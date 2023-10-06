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

#ifndef UNO_GO_HPP
#define UNO_GO_HPP

#pragma once

#include <assert.h>
#include <cstring>
#include <iostream>
#include <ostream>
#include <string>
#include <string_view>
#include <vector>

enum DataType : int32_t {
  kInt64 = 0,
  kFloat32,
  kString,
  kInt64s,
  kFloat32s,
  kStrings,
  kError = 127,
};

// this struct is the same as go string
typedef struct GoString {
  char *ptr;
  size_t len;
  GoString() : ptr(nullptr), len(0) {}
  GoString(size_t size) {
    ptr = new char[size];
    len = size;
    memset(ptr, 0, len);
  }

  GoString(const std::string &str) {
    len = str.size();
    ptr = new char[len];
    memcpy(ptr, str.c_str(), len);
  }

  GoString(const GoString &s) {
    len = s.len;
    ptr = new char[len];
    memcpy(ptr, s.ptr, len);
  }

  void operator=(const GoString &s) {
    if (ptr != nullptr) {
      delete[] ptr;
    }
    len = s.len;
    ptr = new char[len];
    memcpy(ptr, s.ptr, len);
  }

  void operator=(const std::string &s) {
    if (ptr != nullptr) {
      delete[] ptr;
    }
    len = s.size();
    ptr = new char[len];
    memcpy(ptr, s.c_str(), len);
  }

  ~GoString() {
    if (ptr != nullptr) {
      delete[] ptr;
      ptr = nullptr;
      len = 0;
    }
  }

  size_t size() { return len; }

  bool operator==(const GoString &s) {
    if (len != s.len) {
      return false;
    }
    if (ptr == s.ptr) {
      return true;
    }
    for (size_t i = 0; i < len; i++) {
      if (ptr[i] != s.ptr[i]) {
        return false;
      }
    }
    return true;
  }

  bool operator<(const GoString &s) {
    return std::string_view{ptr, len} < std::string_view{s.ptr, s.len};
  }

  bool operator<=(const GoString &s) {
    return std::string_view{ptr, len} <= std::string_view{s.ptr, s.len};
  }

  bool operator>=(const GoString &s) {
    return std::string_view{ptr, len} >= std::string_view{s.ptr, s.len};
  }

  bool operator>(const GoString &s) {
    return std::string_view{ptr, len} > std::string_view{s.ptr, s.len};
  }

  bool operator!=(const GoString &s) {
    return std::string_view{ptr, len} != std::string_view{s.ptr, s.len};
  }

  friend std::ostream &operator<<(std::ostream &os, const GoString &s) {
    os << "\"" << std::string_view{s.ptr, s.len} << "\"";
    return os;
  }

} GoString, *GoStringPtr;

// this struct is the same as go slice
template <typename T> struct GoSlice {
  T *ptr;
  size_t len;
  size_t cap;
  GoSlice() : ptr(new T[8]), len(0), cap(8) { memset(ptr, 0, cap * sizeof(T)); }

  GoSlice(size_t cap) {
    len = 0;
    cap = cap;
    ptr = new T[cap];
    memset(ptr, 0, cap * sizeof(T));
  }

  GoSlice(const GoSlice &slice) {
    len = slice.len;
    cap = slice.cap;
    ptr = new T[cap];
    memcpy(ptr, slice.ptr, len * sizeof(T));
  }

  GoSlice(const std::vector<T> &slice) {
    len = slice.size();
    cap = slice.size();
    ptr = new T[cap];
    for (size_t i = 0; i < len; i++) {
      ptr[i] = slice[i];
    }
  }

  ~GoSlice() {
    if (ptr != nullptr) {
      delete[] ptr;
    }
  }

  void operator=(const GoSlice &slice) {
    if (ptr != nullptr) {
      delete[] ptr;
    }
    len = slice.size();
    cap = slice.size();
    ptr = new T[cap];
    for (size_t i = 0; i < len; i++) {
      ptr[i] = slice[i];
    }
  }

  void operator=(const std::vector<T> &slice) {
    if (ptr != nullptr) {
      delete[] ptr;
    }
    len = slice.size();
    cap = slice.size();
    ptr = new T[cap];
    for (size_t i = 0; i < len; i++) {
      ptr[i] = slice[i];
    }
  }

  T &operator[](size_t i) {
    assert(0 <= i && i < cap);
    return ptr[i];
  }

  void reverse() {
    T v;
    for (size_t i = 0; i < len / 2; i++) {
      v = ptr[i];
      ptr[i] = ptr[len - 1 - i];
      ptr[len - 1 - i] = v;
    }
  }

  void reserve(size_t capacity) {
    cap = capacity;
    T *p = new T[cap];
    memset(p, 0, cap * sizeof(T));

    for (size_t i = 0; i < len; i++) {
      p[i] = ptr[i];
    }
    delete[] ptr;
    ptr = p;
  }

  void resize(size_t capacity) {
    reserve(capacity);
    len = cap;
  }

  void append(const T &v) {
    if (len == cap) {
      reserve(cap * 2);
    }
    ptr[len++] = v;
  }

  void print() {
    std::cout << "[";
    for (size_t i = 0; i < len; i++) {
      if (i) {
        std::cout << ",";
      }
      std::cout << ptr[i];
    }
    std::cout << "]" << std::endl;
  }

  size_t size() { return len; }
};
template <typename T> using Slice = struct GoSlice<T>;

using Int32Slice = Slice<int32_t>;
using Int64Slice = Slice<int64_t>;
using StringSlice = Slice<GoString>;
using Float32Slice = Slice<float>;
using TypeSlice = Slice<DataType>;
using VarSlice = Slice<void *>;

#endif // UNO_GO_HPP