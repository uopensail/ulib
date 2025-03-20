//
// `UNO` - 'uno is expression engine written in c++'
// Copyright (C) 2019 - present timepi <timepi123@gmail.com>
// UNO is provided under: GNU Affero General Public License (AGPL3.0)
// https://www.gnu.org/licenses/agpl-3.0.html unless stated otherwise.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//

#ifndef UNO_FUNC_HPP
#define UNO_FUNC_HPP

#pragma once

#include <functional>
#include <map>
#include <string>

#include "builtins.h"
#include "go.hpp"

struct Function;
using Func = std::function<void *(struct Function *, VarSlice *)>;

struct Function {
  Func func;
  Int32Slice args;
  Function() : func(nullptr) {}
  Function(const Function &f) : func(f.func), args(f.args) {}
  ~Function() {}
  void *operator()(VarSlice *vars) { return func(this, vars); }
};
using Function = struct Function;

template <typename T0, typename... ArgsType, size_t... Is>
T0 *invoke_helper(Function *func, T0 *(*f)(ArgsType *...), VarSlice *vars,
                  std::index_sequence<Is...>) {
  T0 *ret =
      f((static_cast<
          typename std::tuple_element<Is, std::tuple<ArgsType...>>::type *>(
          (*vars)[func->args[Is]]))...);
  return ret;
}

template <typename T0, typename... ArgsType>
T0 *invoke(Function *func, T0 *(*f)(ArgsType *...), VarSlice *vars) {
  for (size_t i = 0; i < sizeof...(ArgsType); i++) {
    if ((*vars)[i] == nullptr) {
      return nullptr;
    }
  }
  return invoke_helper<T0, ArgsType...>(
      func, f, vars, std::make_index_sequence<sizeof...(ArgsType)>{});
}

template <typename T0, typename... ArgsType>
Func get_func(T0 *(*f)(ArgsType *...)) {
  auto foo = [f](Function *func, VarSlice *vars) {
    return invoke<T0, ArgsType...>(func, f, vars);
  };
  return foo;
}

const std::map<std::string, Func> builtin_functions = {
    {"addi", get_func(_add<int64_t>)},
    {"addf", get_func(_add<float>)},
    {"subi", get_func(_sub<int64_t>)},
    {"subf", get_func(_sub<float>)},
    {"muli", get_func(_mul<int64_t>)},
    {"mulf", get_func(_mul<float>)},
    {"divi", get_func(_div<int64_t>)},
    {"divf", get_func(_div<float>)},
    {"mod", get_func(_mod)},
    {"pow", get_func(_pow)},
    {"round", get_func(_round)},
    {"floor", get_func(_floor)},
    {"ceil", get_func(_ceil)},
    {"log", get_func(_log)},
    {"exp", get_func(_exp)},
    {"log10", get_func(_log10)},
    {"log2", get_func(_log2)},
    {"sqrt", get_func(_sqrt)},
    {"abs", get_func(_abs<float>)},
    {"absi", get_func(_abs<int64_t>)},
    {"absf", get_func(_abs<float>)},
    {"sin", get_func(_sin)},
    {"asin", get_func(_asin)},
    {"sinh", get_func(_sinh)},
    {"asinh", get_func(_asinh)},
    {"cos", get_func(_cos)},
    {"acos", get_func(_acos)},
    {"cosh", get_func(_cosh)},
    {"acosh", get_func(_acosh)},
    {"tan", get_func(_tan)},
    {"atan", get_func(_atan)},
    {"tanh", get_func(_tanh)},
    {"atanh", get_func(_atanh)},
    {"sigmoid", get_func(_sigmoid)},
    {"min", get_func(min<float>)},
    {"mini", get_func(min<int64_t>)},
    {"minf", get_func(min<float>)},
    {"max", get_func(max<float>)},
    {"maxi", get_func(max<int64_t>)},
    {"maxf", get_func(max<float>)},
    {"year", get_func(year)},
    {"month", get_func(month)},
    {"day", get_func(day)},
    {"date", get_func(date)},
    {"hour", get_func(hour)},
    {"minute", get_func(minute)},
    {"second", get_func(second)},
    {"now", get_func(now)},
    {"from_unixtime", get_func(from_unixtime)},
    {"unix_timestamp", get_func(unix_timestamp)},
    {"date_add", get_func(date_add)},
    {"date_sub", get_func(date_sub)},
    {"date_diff", get_func(date_diff)},
    {"reverse", get_func(reverse)},
    {"upper", get_func(upper)},
    {"lower", get_func(lower)},
    {"concat", get_func(concat)},
    {"castf2i", get_func(cast<int64_t, float>)},
    {"casti2f", get_func(cast<float, int64_t>)},
};

#endif  // UNO_FUNC_HPP