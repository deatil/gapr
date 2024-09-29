# Copyright 2021 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# FROM coqorg/coq:dev-native-flambda

FROM coqorg/coq:8.20.0-native

RUN git clone https://github.com/mit-plv/fiat-crypto && cd fiat-crypto && \
    git checkout v0.1.4 && \
    git submodule update --init --recursive
RUN cd fiat-crypto && eval $(opam env) && make -j4 standalone-ocaml SKIP_BEDROCK2=1

ENV PATH /home/coq/fiat-crypto/src/ExtractionOCaml:$PATH
