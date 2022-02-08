# Golang BACNet Library

[![CircleCI](https://circleci.com/gh/noahtkeller/go-bacnet/tree/develop.svg?style=shield)](https://circleci.com/gh/noahtkeller/go-bacnet/tree/develop)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bfc4d0c7ef82442eaa6d2c3d18bfb3a7)](https://app.codacy.com/gh/noahtkeller/go-bacnet?utm_source=github.com&utm_medium=referral&utm_content=noahtkeller/go-bacnet&utm_campaign=Badge_Grade_Settings)
[![Join the chat at https://gitter.im/noahtkeller/go-bacnet](https://badges.gitter.im/noahtkeller/go-bacnet.svg)](https://gitter.im/noahtkeller/go-bacnet?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

This is influenced heavily by Steve Karg's bacnet-stack library.
Everything that's here should work, except for the svc_ack_alarm stuff.
Having trouble encoding the strings here. Also had some troubles decoding
signed integers, which is compensated for by using 64 bit integers for now.

### License

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    This permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
