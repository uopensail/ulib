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

#ifndef UNO_H
#define UNO_H

#ifdef __cplusplus
extern "C" {
#endif

void *uno_create_expression(void *str, int len);
void uno_release_expression(void *expression);
int uno_eval(void *expression, char *slice);
void uno_call_unsafe(void *str, int len, char *slice);
void uno_preeval(void *expression, char *slice);
void uno_clean_varslice(void *expression, char *slice);
void uno_batch_eval(void *expression, char *slices, char *result);

#ifdef __cplusplus
};
#endif

#endif  // UNO_H