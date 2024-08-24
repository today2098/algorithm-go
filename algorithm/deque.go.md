---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
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
    RuntimeError: bundler is not specified: algorithm/deque.go\n"
  code: "package algorithm\n\nimport (\n\t\"container/list\"\n\t\"errors\"\n)\n\n\
    var ErrDequeEmpty = errors.New(\"Deque: deque is empty\")\n\ntype Deque[T any]\
    \ struct {\n\tdata *list.List\n}\n\nfunc NewDeque[T any]() *Deque[T] {\n\treturn\
    \ &Deque[T]{data: list.New()}\n}\n\nfunc (dq *Deque[T]) Empty() bool {\n\treturn\
    \ dq.Size() == 0\n}\n\nfunc (dq *Deque[T]) Size() int {\n\treturn dq.data.Len()\n\
    }\n\nfunc (dq *Deque[T]) Front() T {\n\tres := dq.data.Front()\n\tif res == nil\
    \ {\n\t\tpanic(ErrDequeEmpty)\n\t}\n\treturn res.Value.(T)\n}\n\nfunc (dq *Deque[T])\
    \ Back() T {\n\tres := dq.data.Back()\n\tif res == nil {\n\t\tpanic(ErrDequeEmpty)\n\
    \t}\n\treturn res.Value.(T)\n}\n\nfunc (dq *Deque[T]) PushFront(x T) {\n\tdq.data.PushFront(x)\n\
    }\n\nfunc (dq *Deque[T]) PushFrontRange(v []T) {\n\tfor i := len(v) - 1; i >=\
    \ 0; i-- {\n\t\tdq.data.PushFront(v[i])\n\t}\n}\n\nfunc (dq *Deque[T]) PushBack(x\
    \ T) {\n\tdq.data.PushBack(x)\n}\n\nfunc (dq *Deque[T]) PushBackRange(v []T) {\n\
    \tfor i := 0; i < len(v); i++ {\n\t\tdq.data.PushBack(v[i])\n\t}\n}\n\nfunc (dq\
    \ *Deque[T]) PopFront() T {\n\tres := dq.data.Front()\n\tif res == nil {\n\t\t\
    panic(ErrDequeEmpty)\n\t}\n\treturn dq.data.Remove(res).(T)\n}\n\nfunc (dq *Deque[T])\
    \ PopBack() T {\n\tres := dq.data.Back()\n\tif res == nil {\n\t\tpanic(ErrDequeEmpty)\n\
    \t}\n\treturn dq.data.Remove(res).(T)\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/queue.go
  - algorithm/stack.go
  isVerificationFile: false
  path: algorithm/deque.go
  requiredBy:
  - main.go
  - algorithm/queue.go
  - algorithm/stack.go
  timestamp: '2024-08-24 14:32:30+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
documentation_of: algorithm/deque.go
layout: document
redirect_from:
- /library/algorithm/deque.go
- /library/algorithm/deque.go.html
title: algorithm/deque.go
---
