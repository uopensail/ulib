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

#ifndef UNO_EXPR_HPP
#define UNO_EXPR_HPP

#pragma once

#include "node.hpp"

struct Columns {
  Int32Slice addr;
  StringSlice cols;
  TypeSlice types;

  Columns() {}
  ~Columns() {}

  void reserve(size_t cap) {
    addr.reserve(cap);
    cols.reserve(cap);
    types.reserve(cap);
  }

  void append(int32_t id, const std::string &str, DataType type) {
    addr.append(id);
    cols.append(str);
    types.append(type);
  }
};

struct Expression {
  Columns columns;
  VarSlice varslice;
  NodePtrSlice nodes;
  Expression(const std::string_view &str) {
    json doc = json::parse(str);
    const std::vector<json> &array = doc["nodes"].get<std::vector<json>>();
    size_t size = array.size();
    columns.reserve(size);

    // create node slice
    nodes.resize(size);

    // create varslice
    varslice.resize(size);

    for (size_t i = 0; i < size; i++) {
      NodeType ntype = array[i]["ntype"].get<NodeType>();
      int32_t id = array[i]["id"].get<int32_t>();
      if (ntype == NodeType::kLiteralNode) {
        nodes[id] = new LiteralNode(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kVarNode) {
        VarNode *n = new VarNode(array[i]);
        nodes[id] = n;
        columns.append(id, n->value, n->dtype);
      } else if (ntype == NodeType::kInt64Node) {
        nodes[id] = new Int64Node(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kInt64SliceNode) {
        nodes[id] = new Int64SliceNode(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kFloat32Node) {
        nodes[id] = new Float32Node(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kFloat32SliceNode) {
        nodes[id] = new Float32SliceNode(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kStringNode) {
        nodes[id] = new StringNode(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kStringSliceNode) {
        nodes[id] = new StringSliceNode(array[i]);
        (*(nodes[id]))(&varslice);
      } else if (ntype == NodeType::kFunctionNode) {
        nodes[id] = new FunctionNode(array[i]);
      } else if (ntype == NodeType::kAndNode) {
        nodes[id] = new AndNode(array[i]);
      } else if (ntype == NodeType::kOrNode) {
        nodes[id] = new OrNode(array[i]);
      } else if (ntype == NodeType::kCmpNode) {
        nodes[id] = new CmpNode(array[i]);
      } else if (ntype == NodeType::kInNode) {
        nodes[id] = new InNode(array[i]);
      } else if (ntype == NodeType::kNotInNode) {
        nodes[id] = new NotInNode(array[i]);
      }
    }
  }

  ~Expression() {
    for (size_t i = 0; i < nodes.size(); i++) {
      if (nodes[i] != nullptr) {
        delete nodes[i];
        nodes[i] = nullptr;
      }
    }

    // release varslice
    for (size_t i = 0; i < varslice.size(); i++) {
      if (varslice[i] != nullptr) {
        free(varslice[i]);
        varslice[i] = nullptr;
      }
    }
  }

  int32_t operator()(VarSlice *vars) {
    for (size_t i = 0; i < nodes.size(); i++) {
      Node *ptr = nodes[i];
      if (ptr == nullptr) {
        return -1;
      }
      (*ptr)(vars);
    }
    void *ptr = (*vars)[vars->size() - 1];
    if (ptr != nullptr) {
      if (*(bool *)(ptr)) {
        return 1;
      }
      return 0;
    }
    return -1;
  }
};

#endif // UNO_EXPR_HPP