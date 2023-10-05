#include "builtins.h"

int64_t *_mod(int64_t *a, int64_t *b) {
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  assert(*b != 0);
  *ret = (*a) % (*b);
  return ret;
}

float *_pow(float *a, float *b) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = powf(*a, *b);
  return ret;
}

int64_t *_round(float *x) {
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  *ret = int64_t(roundf(*x));
  return ret;
}

int64_t *_floor(float *x) {
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  *ret = int64_t(floorf(*x));
  return ret;
}

int64_t *_ceil(float *x) {
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  *ret = int64_t(ceilf(*x));
  return ret;
}

float *_log(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = logf(*x);
  return ret;
}

float *_exp(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = expf(*x);
  return ret;
}

float *_log10(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = log10(*x);
  return ret;
}

float *_log2(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = log2f(*x);
  return ret;
}

float *_sqrt(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = sqrtf(*x);
  return ret;
}

float *_sin(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = sinf(*x);
  return ret;
}

float *_asin(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = asinf(*x);
  return ret;
}

float *_sinh(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = sinhf(*x);
  return ret;
}

float *_asinh(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = asinhf(*x);
  return ret;
}

float *_cos(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = cosf(*x);
  return ret;
}

float *_acos(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = acosf(*x);
  return ret;
}

float *_cosh(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = coshf(*x);
  return ret;
}

float *_acosh(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = acoshf(*x);
  return ret;
}

float *_tan(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = tanf(*x);
  return ret;
}

float *_atan(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = atanf(*x);
  return ret;
}

float *_tanh(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = tanhf(*x);
  return ret;
}

float *_atanh(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = atanhf(*x);
  return ret;
}

float *_sigmoid(float *x) {
  float *ret = (float *)malloc(sizeof(float));
  *ret = 1.0 / (1.0 + expf(-(*x)));
  return ret;
}

GoStringPtr year() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%Y", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

GoStringPtr month() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%m", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

GoStringPtr day() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%d", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

GoStringPtr date() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%Y%m%d", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

GoStringPtr hour() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%H", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

GoStringPtr minute() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%M", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

GoStringPtr second() {
  time_t t = time(nullptr);
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];
  strftime(ptr, 64, "%S", localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

int64_t *now() {
  auto now = std::chrono::duration_cast<std::chrono::milliseconds>(
      std::chrono::system_clock::now().time_since_epoch());
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  *ret = now.count();
  return ret;
}

GoStringPtr from_unixtime(int64_t *ts, GoStringPtr format) {
  char *tmp = (char *)calloc(sizeof(GoString) + 64, 1);
  GoStringPtr ret = (GoStringPtr)tmp;
  char *ptr = &ptr[sizeof(GoString)];

  time_t t = *ts;
  strftime(tmp, 64, format->ptr, localtime(&t));
  ret->len = strlen(ptr);
  return ret;
}

int64_t *unix_timestamp(GoStringPtr t, GoStringPtr format) {
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  std::tm timeinfo = {};
  strptime(t->ptr, format->ptr, &timeinfo);
  *ret = mktime(&timeinfo);
  return ret;
}

GoStringPtr date_add(GoStringPtr s, int64_t *n) {
  GoString format("%Y-%m-%d");
  int64_t *ts = unix_timestamp(s, &format);
  *ts += 86400 * (*n);
  GoStringPtr ret = from_unixtime(ts, &format);
  free(ts);
  return ret;
}

GoStringPtr date_sub(GoStringPtr s, int64_t *n) {
  int64_t v = -(*n);
  return date_add(s, &v);
}

int64_t *date_diff(GoStringPtr d1, GoStringPtr d2) {
  GoString format("%Y-%m-%d");
  int64_t *t1 = unix_timestamp(d1, &format);
  int64_t *t2 = unix_timestamp(d2, &format);
  int64_t *ret = (int64_t *)malloc(sizeof(int64_t));
  *ret = int64_t(((*t1) - (*t2)) / 86400);
  free(t1);
  free(t2);
  return ret;
}

GoStringPtr reverse(GoStringPtr s) {
  GoStringPtr ret = (GoStringPtr)malloc(sizeof(GoString) + s->len);
  ret->ptr = (char *)(ret) + sizeof(GoString);
  ret->len = s->len;
  for (size_t i = 0; i < s->len; i++) {
    ret->ptr[s->len - 1 - i] = s->ptr[i];
  }
  return ret;
}

GoStringPtr upper(GoStringPtr s) {
  GoStringPtr ret = (GoStringPtr)malloc(sizeof(GoString) + s->len);
  ret->ptr = (char *)(ret) + sizeof(GoString);
  ret->len = s->len;
  for (size_t i = 0; i < s->len; i++) {
    ret->ptr[i] = toupper(s->ptr[i]);
  }
  return ret;
}

GoStringPtr lower(GoStringPtr s) {
  GoStringPtr ret = (GoStringPtr)malloc(sizeof(GoString) + s->len);
  ret->ptr = (char *)(ret) + sizeof(GoString);
  ret->len = s->len;
  for (size_t i = 0; i < s->len; i++) {
    ret->ptr[i] = tolower(s->ptr[i]);
  }
  return ret;
}

GoStringPtr substr(GoStringPtr s, int64_t *start, int64_t *len) {
  GoStringPtr ret = (GoStringPtr)malloc(sizeof(GoString) + (*len));
  ret->len = *len;
  ret->ptr = (char *)(ret) + sizeof(GoString);
  memcpy(ret->ptr, &s->ptr[*start], *len);
  return ret;
}

GoStringPtr concat(GoStringPtr a, GoStringPtr b) {
  GoStringPtr ret = (GoStringPtr)malloc(sizeof(GoString) + a->len + b->len);
  ret->len = a->len + b->len;
  ret->ptr = (char *)(ret) + sizeof(GoString);
  memcpy(ret->ptr, a->ptr, a->len);
  memcpy(ret->ptr + a->len, b->ptr, b->len);
  return ret;
}