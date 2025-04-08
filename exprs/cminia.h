//
// `Minia` - A C++ tool for feature transformation and hashing
// Copyright (C) 2019 - present Uopensail <timepi123@gmail.com>
// This software is distributed under the GNU Affero General Public License
// (AGPL3.0) For more information, please visit:
// https://www.gnu.org/licenses/agpl-3.0.html
//
// This program is free software: you are free to redistribute and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. Refer to the
// GNU Affero General Public License for more details.
//

#ifndef C_MINIA_H_
#define C_MINIA_H_
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

/**
 * @brief Creates a new Minia instance
 * @param config_path Path to configuration file
 * @return Pointer to Minia instance or nullptr on failure
 */
void *minia_create(const char *config_path);

/**
 * @brief Releases a Minia instance
 * @param m Pointer to Minia instance
 */
void minia_release(void *m);

/**
 * @brief Get all feature keys
 * @param m Pointer to Minia instance
 * @return Pointer to Minia keys
 */
void *minia_features(void *m);

/**
 * @brief Executes Minia processing
 * @param m Minia instance pointer
 * @param data Features data string, JSON format
 * @return Pointer to processed features or nullptr on failure
 */
void *minia_call(void *m, const char *data);

/**
 * @brief Retrieves a feature value
 * @param features Features container pointer
 * @param key Feature key name
 * @param type Output parameter for data type
 * @return Pointer to feature data or nullptr
 */
void *minia_get_feature(void *features, const char *key, int32_t *type);

/**
 * @brief Releases feature memory
 * @param feature Feature data pointer
 * @param type Data type identifier
 */
void minia_del_feature(void *feature, int32_t type);

/**
 * @brief Releases features container
 * @param features Features container pointer
 */
void minia_del_features(void *features);

#ifdef __cplusplus
} // extern "C"

#endif

#endif // C_MINIA_H_