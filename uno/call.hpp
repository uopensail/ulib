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

#ifndef UNO_CALL_HPP
#define UNO_CALL_HPP

#pragma once

#include <functional>
#include <map>
#include <string>

#include "builtins.h"
#include "go.hpp"

using Call = std::function<void(VarSlice *)>;

template <typename T0, typename... ArgsType, size_t... Is>
T0 *invoke_helper(T0 *(*f)(ArgsType *...), VarSlice *vars,
                  std::index_sequence<Is...>) {
  T0 *ret =
      f((static_cast<
          typename std::tuple_element<Is, std::tuple<ArgsType...>>::type *>(
          (*vars)[Is]))...);
  return ret;
}

template <typename T0, typename... ArgsType>
T0 *invoke(T0 *(*f)(ArgsType *...), VarSlice *vars) {
  for (size_t i = 0; i < sizeof...(ArgsType); i++) {
    if ((*vars)[i] == nullptr) {
      return nullptr;
    }
  }
  return invoke_helper<T0, ArgsType...>(
      f, vars, std::make_index_sequence<sizeof...(ArgsType)>{});
}

template <typename T0, typename... ArgsType>
Call get_call(T0 *(*f)(ArgsType *...)) {
  auto foo = [f](VarSlice *vars) {
    T0 *ret = invoke<T0, ArgsType...>(f, vars);
    vars->ptr[vars->len - 1] = ret;
  };
  return foo;
}

const std::map<std::string, Call> builtin_callers = {
    {"addi", get_call(_add<int64_t>)},
    {"addf", get_call(_add<float>)},
    {"subi", get_call(_sub<int64_t>)},
    {"subf", get_call(_sub<float>)},
    {"muli", get_call(_mul<int64_t>)},
    {"mulf", get_call(_mul<float>)},
    {"divi", get_call(_div<int64_t>)},
    {"divf", get_call(_div<float>)},
    {"mod", get_call(_mod)},
    {"pow", get_call(_pow)},
    {"round", get_call(_round)},
    {"floor", get_call(_floor)},
    {"ceil", get_call(_ceil)},
    {"log", get_call(_log)},
    {"exp", get_call(_exp)},
    {"log10", get_call(_log10)},
    {"log2", get_call(_log2)},
    {"sqrt", get_call(_sqrt)},
    {"abs", get_call(_abs<float>)},
    {"absi", get_call(_abs<int64_t>)},
    {"absf", get_call(_abs<float>)},
    {"sin", get_call(_sin)},
    {"asin", get_call(_asin)},
    {"sinh", get_call(_sinh)},
    {"asinh", get_call(_asinh)},
    {"cos", get_call(_cos)},
    {"acos", get_call(_acos)},
    {"cosh", get_call(_cosh)},
    {"acosh", get_call(_acosh)},
    {"tan", get_call(_tan)},
    {"atan", get_call(_atan)},
    {"tanh", get_call(_tanh)},
    {"atanh", get_call(_atanh)},
    {"sigmoid", get_call(_sigmoid)},
    {"min", get_call(min<float>)},
    {"mini", get_call(min<int64_t>)},
    {"minf", get_call(min<float>)},
    {"max", get_call(max<float>)},
    {"maxi", get_call(max<int64_t>)},
    {"maxf", get_call(max<float>)},
    {"year", get_call(year)},
    {"month", get_call(month)},
    {"day", get_call(day)},
    {"date", get_call(date)},
    {"hour", get_call(hour)},
    {"minute", get_call(minute)},
    {"second", get_call(second)},
    {"now", get_call(now)},
    {"from_unixtime", get_call(from_unixtime)},
    {"unix_timestamp", get_call(unix_timestamp)},
    {"date_add", get_call(date_add)},
    {"date_sub", get_call(date_sub)},
    {"date_diff", get_call(date_diff)},
    {"reverse", get_call(reverse)},
    {"upper", get_call(upper)},
    {"lower", get_call(lower)},
    {"concat", get_call(concat)},
    {"castf2i", get_call(cast<int64_t, float>)},
    {"casti2f", get_call(cast<float, int64_t>)},
};

#endif  // UNO_FUNC_HPP