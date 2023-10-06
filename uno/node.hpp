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

#ifndef UNO_NODE_HPP
#define UNO_NODE_HPP

#pragma once

#include "func.hpp"
#include "json.hpp"

using json = nlohmann::json;

// this file defines all types of nodes

enum NodeType : int32_t {
  kVarNode = 0,
  kInt64Node,
  kInt64SliceNode,
  kFloat32Node,
  kFloat32SliceNode,
  kStringNode,
  kStringSliceNode,
  kFunctionNode,
  kAndNode,
  kOrNode,
  kCmpNode,
  kInNode,
  kNotInNode,
  kLiteralNode,
  kErrorNode = 127,
};

enum CmpType : int32_t {
  kNil = 0,
  kEqual,
  kNotEqual,
  kGreaterThan,
  kGreaterThanEqual,
  kLessThan,
  kLessThanEqual,
};

struct Node {
  int32_t id;
  Node(int32_t i) : id(i) {}
  virtual ~Node() {}
  virtual NodeType type() { return NodeType::kErrorNode; }
  virtual void operator()(VarSlice *) {}
};

using NodePtrSlice = Slice<Node *>;

struct VarNode : Node {
  std::string value;
  DataType dtype;
  VarNode(int32_t i, const std::string &str, DataType type)
      : Node(i), value(str), dtype(type) {}
  VarNode(const json &json)
      : VarNode(json["id"].get<int32_t>(), json["value"].get<std::string>(),
                json["dtype"].get<DataType>()) {}
  virtual ~VarNode() {}
  virtual NodeType type() { return NodeType::kVarNode; }
  virtual void operator()(VarSlice *vars) {}
};

struct LiteralNode : Node {
  bool value;
  LiteralNode(int32_t i, bool v) : Node(i), value(v) {}
  LiteralNode(const json &json)
      : LiteralNode(json["id"].get<int32_t>(), json["value"].get<bool>()) {}
  virtual ~LiteralNode() {}
  virtual NodeType type() { return NodeType::kLiteralNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }
    bool *v = (bool *)malloc(sizeof(bool));
    *v = value;
    (*vars)[id] = v;
  }
};

struct Int64Node : Node {
  int64_t value;
  Int64Node(int32_t i, int64_t v) : Node(i), value(v) {}
  Int64Node(const json &json)
      : Int64Node(json["id"].get<int32_t>(), json["value"].get<int64_t>()) {}

  virtual ~Int64Node() {}
  virtual NodeType type() { return NodeType::kInt64Node; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    int64_t *v = (int64_t *)malloc(sizeof(int64_t));
    *v = value;
    (*vars)[id] = v;
  }
};

struct Int64SliceNode : Node {
  Int64Slice value;
  Int64SliceNode(int32_t i, const Int64Slice &v) : Node(i), value(v) {}
  Int64SliceNode(const json &json)
      : Int64SliceNode(json["id"].get<int32_t>(),
                       Int64Slice{json["value"].get<std::vector<int64_t>>()}) {}
  virtual ~Int64SliceNode() {}
  virtual NodeType type() { return NodeType::kInt64SliceNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }
    Int64Slice *v = (Int64Slice *)malloc(sizeof(Int64Slice));
    v->cap = value.cap;
    v->len = value.len;
    v->ptr = value.ptr;
    (*vars)[id] = v;
  }
};

struct Float32Node : Node {
  float value;
  Float32Node(int32_t i, float v) : Node(i), value(v) {}
  Float32Node(const json &json)
      : Float32Node(json["id"].get<int32_t>(), json["value"].get<float>()) {}
  virtual ~Float32Node() {}
  virtual NodeType type() { return NodeType::kFloat32Node; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    float *v = (float *)malloc(sizeof(float));
    *v = value;
    (*vars)[id] = v;
  }
};

struct Float32SliceNode : Node {
  Float32Slice value;
  Float32SliceNode(int32_t i, const Float32Slice &v) : Node(i), value(v) {}
  Float32SliceNode(const json &json)
      : Float32SliceNode(
            json["id"].get<int32_t>(),
            Float32Slice{json["value"].get<std::vector<float>>()}) {}

  virtual ~Float32SliceNode() {}
  virtual NodeType type() { return NodeType::kFloat32SliceNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    Float32Slice *v = (Float32Slice *)malloc(sizeof(Float32Slice));
    v->cap = value.cap;
    v->len = value.len;
    v->ptr = value.ptr;
    (*vars)[id] = v;
  }
};

struct StringNode : Node {
  std::string value;
  StringNode(int32_t i, const std::string &v) : Node(i), value(v) {}
  StringNode(const json &json)
      : StringNode(json["id"].get<int32_t>(),
                   json["value"].get<std::string>()) {}
  virtual ~StringNode() {}
  virtual NodeType type() { return NodeType::kStringNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    GoString *v = (GoString *)malloc(sizeof(GoString));
    v->len = value.size();
    v->ptr = (char *)value.c_str();
    (*vars)[id] = v;
  }
};

struct StringSliceNode : Node {
  StringSlice value;
  StringSliceNode(int32_t i, StringSlice &v) : Node(i), value(v) {}
  StringSliceNode(const json &json) : Node(json["id"].get<int32_t>()) {
    const std::vector<std::string> &slice =
        json["value"].get<std::vector<std::string>>();
    value.reserve(slice.size());
    for (size_t i = 0; i < slice.size(); i++) {
      value.append(slice[i]);
    }
  }

  virtual ~StringSliceNode() {}

  virtual NodeType type() { return NodeType::kStringSliceNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    StringSlice *v = (StringSlice *)malloc(sizeof(StringSlice));
    v->len = value.len;
    v->ptr = value.ptr;
    v->cap = value.cap;
    (*vars)[id] = v;
  }
};

struct FunctionNode : Node {
  Function func;
  FunctionNode(int32_t i, Function &v) : Node(i), func(v) {}
  FunctionNode(const json &json) : Node(json["id"].get<int32_t>()) {
    func.args = json["args"].get<std::vector<int32_t>>();
    auto iter = builtin_functions.find(json["func"].get<std::string>());
    assert(iter != builtin_functions.end());
    func.func = iter->second;

#if !defined(__clang__)
    func.args.reverse();
#endif
  }

  virtual ~FunctionNode() {}

  virtual NodeType type() { return NodeType::kFunctionNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }
    (*vars)[id] = func(vars);
  }
};

struct AndNode : Node {
  int32_t left;
  int32_t right;
  AndNode(int32_t i, int32_t l, int32_t r) : Node(i), left(l), right(r) {}

  AndNode(const json &json)
      : AndNode(json["id"].get<int32_t>(), json["left"].get<int32_t>(),
                json["right"].get<int32_t>()) {}

  virtual ~AndNode() {}

  virtual NodeType type() { return NodeType::kAndNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }
    if ((*vars)[left] == nullptr && (*vars)[right] == nullptr) {
      (*vars)[id] = nullptr;
      return;
    }

    if ((*vars)[left] != nullptr && !(*static_cast<bool *>((*vars)[left]))) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = false;
      (*vars)[id] = ret;
      return;
    }
    if ((*vars)[right] != nullptr && !(*static_cast<bool *>((*vars)[right]))) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = false;
      (*vars)[id] = ret;
      return;
    }
    bool *ret = (bool *)malloc(sizeof(bool));
    *ret = true;
    (*vars)[id] = ret;
  }
};

struct OrNode : Node {
  int32_t left;
  int32_t right;
  OrNode(int32_t i, int32_t l, int32_t r) : Node(i), left(l), right(r) {}
  OrNode(const json &json)
      : OrNode(json["id"].get<int32_t>(), json["left"].get<int32_t>(),
               json["right"].get<int32_t>()) {}
  virtual ~OrNode() {}
  virtual NodeType type() { return NodeType::kOrNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    void *lv = (*vars)[left];
    void *rv = (*vars)[right];

    if (lv == nullptr && rv == nullptr) {
      (*vars)[id] = nullptr;
      return;
    }
    if (lv != nullptr && *static_cast<bool *>(lv)) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = true;
      (*vars)[id] = ret;
      return;
    }
    if (rv != nullptr && *static_cast<bool *>(rv)) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = true;
      (*vars)[id] = ret;
      return;
    }
    bool *ret = (bool *)malloc(sizeof(bool));
    *ret = false;
    (*vars)[id] = ret;
  }
};

template <typename T> static bool Compare(CmpType op, T *a, T *b) {
  switch (op) {
  case kEqual:
    return *a == *b;
  case kNotEqual:
    return *a != *b;
  case kGreaterThan:
    return *a > *b;
  case kGreaterThanEqual:
    return *a >= *b;
  case kLessThan:
    return *a < *b;
  case kLessThanEqual:
    return *a <= *b;
  default:
    return false;
  }
}

struct CmpNode : Node {
  int32_t left;
  int32_t right;
  DataType dtype;
  CmpType op;
  CmpNode(int32_t i, int32_t l, int32_t r, DataType datatype, CmpType cmp)
      : Node(i), left(l), right(r), dtype(datatype), op(cmp) {}
  CmpNode(const json &json)
      : CmpNode(json["id"].get<int32_t>(), json["left"].get<int32_t>(),
                json["right"].get<int32_t>(), json["dtype"].get<DataType>(),
                json["cmp"].get<CmpType>()) {}
  virtual ~CmpNode() {}

  virtual NodeType type() { return NodeType::kCmpNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }

    void *lv = (*vars)[left];
    void *rv = (*vars)[right];
    if (lv == nullptr || rv == nullptr) {
      return;
    }
    if (dtype == DataType::kInt64) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = Compare<int64_t>(op, static_cast<int64_t *>(lv),
                              static_cast<int64_t *>(rv));
      (*vars)[id] = ret;
    } else if (dtype == DataType::kFloat32) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = Compare<float>(op, static_cast<float *>(lv),
                            static_cast<float *>(rv));
      (*vars)[id] = ret;
    } else if (dtype == DataType::kString) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = Compare<GoString>(op, static_cast<GoString *>(lv),
                               static_cast<GoString *>(rv));
      (*vars)[id] = ret;
    }
  }
};

template <typename T> static bool InSlice(T &v, Slice<T> &slice) {
  for (size_t i = 0; i < slice.size(); i++) {
    if (v == slice[i]) {
      return true;
    }
  }
  return false;
}

template <typename T> static bool NotInSlice(T &v, Slice<T> &slice) {
  for (size_t i = 0; i < slice.size(); i++) {
    if (v == slice[i]) {
      return false;
    }
  }
  return true;
}

struct InNode : Node {
  int32_t left;
  int32_t right;
  DataType dtype;

  InNode(int32_t i, int32_t l, int32_t r, DataType datatype)
      : Node(i), left(l), right(r), dtype(datatype) {}
  InNode(const json &json)
      : InNode(json["id"].get<int32_t>(), json["left"].get<int32_t>(),
               json["right"].get<int32_t>(), json["dtype"].get<DataType>()) {}
  virtual ~InNode() {}

  virtual NodeType type() { return NodeType::kInNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }
    void *lv = (*vars)[left];
    void *rv = (*vars)[right];

    if (lv == nullptr || rv == nullptr) {
      return;
    }

    if (dtype == DataType::kInt64) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = InSlice<int64_t>(*static_cast<int64_t *>(lv),
                              *static_cast<Int64Slice *>(rv));
      (*vars)[id] = ret;
    } else if (dtype == DataType::kFloat32) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = InSlice<float>(*static_cast<float *>(lv),
                            *static_cast<Float32Slice *>((*vars)[right]));
      (*vars)[id] = ret;
    } else if (dtype == DataType::kString) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = InSlice<GoString>(*static_cast<GoString *>((*vars)[left]),
                               *static_cast<StringSlice *>((*vars)[right]));
      (*vars)[id] = ret;
    }
  }
};

struct NotInNode : Node {
  int32_t left;
  int32_t right;
  DataType dtype;
  NotInNode(int32_t i, int32_t l, int32_t r, DataType datatype)
      : Node(i), left(l), right(r), dtype(datatype) {}
  NotInNode(const json &json)
      : NotInNode(json["id"].get<int32_t>(), json["left"].get<int32_t>(),
                  json["right"].get<int32_t>(), json["dtype"].get<DataType>()) {
  }
  virtual ~NotInNode() {}
  virtual NodeType type() { return NodeType::kNotInNode; }
  virtual void operator()(VarSlice *vars) {
    if ((*vars)[id] != nullptr) {
      return;
    }
    if (dtype == DataType::kInt64) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = NotInSlice<int64_t>(*static_cast<int64_t *>((*vars)[left]),
                                 *static_cast<Int64Slice *>((*vars)[right]));
      (*vars)[id] = ret;
    } else if (dtype == DataType::kFloat32) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = NotInSlice<float>(*static_cast<float *>((*vars)[left]),
                               *static_cast<Float32Slice *>((*vars)[right]));
      (*vars)[id] = ret;
    } else if (dtype == DataType::kString) {
      bool *ret = (bool *)malloc(sizeof(bool));
      *ret = NotInSlice<GoString>(*static_cast<GoString *>((*vars)[left]),
                                  *static_cast<StringSlice *>((*vars)[right]));
      (*vars)[id] = ret;
    }
  }
};

#endif // UNO_NODE_HPP