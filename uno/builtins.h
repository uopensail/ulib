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

#ifndef UNO_BUILTINS_H
#define UNO_BUILTINS_H

#pragma once

#include <assert.h>
#include <math.h>
#include <time.h>

#include <chrono>
#include <iostream>
#include <stdexcept>
#include <string>
#include <type_traits>
#include <vector>

#include "go.hpp"

template <typename T>
T *_add(T *a, T *b) {
  T *ret = (T *)malloc(sizeof(T));
  *ret = *a + *b;
  return ret;
}

template <typename T>
T *_sub(T *a, T *b) {
  T *ret = (T *)malloc(sizeof(T));
  *ret = *a - *b;
  return ret;
}

template <typename T>
T *_mul(T *a, T *b) {
  T *ret = (T *)malloc(sizeof(T));
  *ret = (*a) * (*b);
  return ret;
}

template <typename T>
T *_div(T *a, T *b) {
  assert(*b != 0);
  T *ret = (T *)malloc(sizeof(T));
  *ret = (*a) / (*b);
  return ret;
}

int64_t *_mod(int64_t *a, int64_t *b);
float *_pow(float *a, float *b);
int64_t *_round(float *x);
int64_t *_floor(float *x);
int64_t *_ceil(float *x);
float *_log(float *x);
float *_exp(float *x);
float *_log10(float *x);
float *_log2(float *x);
float *_sqrt(float *x);

template <typename T>
T *_abs(T *x) {
  T *ret = (T *)malloc(sizeof(T));
  if ((*x) < 0) {
    *ret = -(*x);
  } else {
    *ret = *x;
  }
  return ret;
}

float *_sin(float *x);
float *_asin(float *x);
float *_sinh(float *x);
float *_asinh(float *x);
float *_cos(float *x);
float *_acos(float *x);
float *_cosh(float *x);
float *_acosh(float *x);
float *_tan(float *x);
float *_atan(float *x);
float *_tanh(float *x);
float *_atanh(float *x);
float *_sigmoid(float *x);

template <typename T>
T *min(Slice<T> *src) {
  assert(src->len > 0);
  T *ret = (T *)malloc(sizeof(T));
  *ret = (*src)[0];
  for (size_t i = 1; i < src->len; i++) {
    if ((*src)[i] < *ret) {
      *ret = (*src)[i];
    }
  }
  return ret;
}

template <typename T>
T *max(Slice<T> *src) {
  assert(src->len > 0);
  T *ret = (T *)malloc(sizeof(T));
  *ret = (*src)[0];
  for (size_t i = 1; i < src->len; i++) {
    if ((*src)[i] > *ret) {
      *ret = (*src)[i];
    }
  }
  return ret;
}

template <typename T0, typename T1>
T0 *cast(T1 *value) {
  if ((!std::is_same_v<T0, int64_t> && !std::is_same_v<T0, float>) ||
      (!std::is_same_v<T1, int64_t> && !std::is_same_v<T1, float>)) {
    throw std::runtime_error("type cast only supports int64_t and float");
  }

  if constexpr (std::is_same_v<T0, T1>) {
    T0 *ret = (T0 *)malloc(sizeof(T0));
    *ret = *value;
    return ret;
  }

  T0 *ret = (T0 *)malloc(sizeof(T0));
  *ret = static_cast<T0>(*value);
  return ret;
}

static bool contain_nullptr() { return false; }

template <typename T, typename... Args>
static bool contain_nullptr(T *arg, Args *...args) {
  if constexpr (sizeof...(Args) == 0) {
    return arg == nullptr;
  } else {
    return contain_nullptr(args...);
  }
}

template <typename T0, typename... ArgsType>
T0 *safe_func_call(T0 *(*func)(ArgsType *...), ArgsType *...args) {
  if (contain_nullptr(args...)) {
    return nullptr;
  }
  return func(args...);
}

GoStringPtr year();
GoStringPtr month();
GoStringPtr day();
GoStringPtr date();
GoStringPtr hour();
GoStringPtr minute();
GoStringPtr second();
int64_t *now();
GoStringPtr from_unixtime(int64_t *ts, GoStringPtr format);
int64_t *unix_timestamp(GoStringPtr t, GoStringPtr format);
GoStringPtr date_add(GoStringPtr s, int64_t *n);
GoStringPtr date_sub(GoStringPtr s, int64_t *n);
int64_t *date_diff(GoStringPtr d1, GoStringPtr d2);
GoStringPtr reverse(GoStringPtr s);
GoStringPtr upper(GoStringPtr s);
GoStringPtr lower(GoStringPtr s);
GoStringPtr substr(GoStringPtr s, int64_t *start, int64_t *len);
GoStringPtr concat(GoStringPtr a, GoStringPtr b);

#endif  // UNO_BUILTINS_H