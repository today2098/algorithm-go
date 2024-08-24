---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/aoj-ALDS1_3_C.test.go
    title: algorithm/deque/aoj-ALDS1_3_C.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/stack.go
    title: algorithm/stack/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/stack.go
    title: algorithm/stack/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/aoj-ALDS1_3_C.test.go
    title: algorithm/deque/aoj-ALDS1_3_C.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: algorithm/deque/deque.go\n"
  code: "package deque\n\nimport (\n\t\"container/list\"\n\t\"errors\"\n)\n\nvar ErrDequeEmpty\
    \ = errors.New(\"Deque: deque is empty\")\n\ntype Deque[T any] struct {\n\tData\
    \ *list.List\n}\n\nfunc NewDeque[T any]() *Deque[T] {\n\treturn &Deque[T]{Data:\
    \ list.New()}\n}\n\nfunc (dq *Deque[T]) Empty() bool {\n\treturn dq.Size() ==\
    \ 0\n}\n\nfunc (dq *Deque[T]) Size() int {\n\treturn dq.Data.Len()\n}\n\nfunc\
    \ (dq *Deque[T]) Front() T {\n\tres := dq.Data.Front()\n\tif res == nil {\n\t\t\
    panic(ErrDequeEmpty)\n\t}\n\treturn res.Value.(T)\n}\n\nfunc (dq *Deque[T]) Back()\
    \ T {\n\tres := dq.Data.Back()\n\tif res == nil {\n\t\tpanic(ErrDequeEmpty)\n\t\
    }\n\treturn res.Value.(T)\n}\n\nfunc (dq *Deque[T]) PushFront(x T) {\n\tdq.Data.PushFront(x)\n\
    }\n\nfunc (dq *Deque[T]) PushFrontRange(v []T) {\n\tfor i := len(v) - 1; i >=\
    \ 0; i-- {\n\t\tdq.Data.PushFront(v[i])\n\t}\n}\n\nfunc (dq *Deque[T]) PushBack(x\
    \ T) {\n\tdq.Data.PushBack(x)\n}\n\nfunc (dq *Deque[T]) PushBackRange(v []T) {\n\
    \tfor i := 0; i < len(v); i++ {\n\t\tdq.Data.PushBack(v[i])\n\t}\n}\n\nfunc (dq\
    \ *Deque[T]) PopFront() T {\n\tres := dq.Data.Front()\n\tif res == nil {\n\t\t\
    panic(ErrDequeEmpty)\n\t}\n\treturn dq.Data.Remove(res).(T)\n}\n\nfunc (dq *Deque[T])\
    \ PopBack() T {\n\tres := dq.Data.Back()\n\tif res == nil {\n\t\tpanic(ErrDequeEmpty)\n\
    \t}\n\treturn dq.Data.Remove(res).(T)\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/stack/stack.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/queue/queue.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
  isVerificationFile: false
  path: algorithm/deque/deque.go
  requiredBy:
  - main.go
  - algorithm/stack/stack.go
  - algorithm/queue/queue.go
  timestamp: '2024-08-24 17:41:38+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
documentation_of: algorithm/deque/deque.go
layout: document
redirect_from:
- /library/algorithm/deque/deque.go
- /library/algorithm/deque/deque.go.html
title: algorithm/deque/deque.go
---
