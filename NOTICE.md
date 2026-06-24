# NOTICE

**Craftgate Go Client** (`github.com/craftgate/craftgate-go-client`)

Copyright (c) 2022 Craftgate

This product is developed and maintained by Craftgate (https://craftgate.io).

This software is licensed under the terms of the Apache License, Version 2.0
(the "License"); you may not use this software except in compliance with the
License. A full copy of the License is included in the [`LICENSE`](./LICENSE)
file distributed with this software.

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

---

## Third-Party Software

This product depends on the third-party software components listed below. Each
component is the property of its respective owners and is governed by its own
license. The Craftgate Go Client does not modify these components; they are
referenced as external modules and resolved via Go modules.

License texts are available at the URLs provided. Inclusion of a component in
this list does not imply endorsement by its authors.

### Runtime Dependencies

These modules are required at runtime and are compiled into binaries that
import the Craftgate Go Client.

#### gorilla/schema

- **Module:** `github.com/gorilla/schema@v1.4.1`
- **Copyright:** Copyright (c) 2012 The Gorilla Authors
- **License:** BSD 3-Clause "New" or "Revised" License
- **Homepage:** https://github.com/gorilla/schema
- **License text:** https://github.com/gorilla/schema/blob/main/LICENSE

### Test-Only Dependencies

These modules are used by the test suite only. They are **not** required by
consumers of the Craftgate Go Client at runtime.

#### testify

- **Module:** `github.com/stretchr/testify@v1.11.1`
- **Copyright:** Copyright (c) 2012-2020 Mat Ryer, Tyler Bunnell and contributors
- **License:** MIT License
- **Homepage:** https://github.com/stretchr/testify
- **License text:** https://github.com/stretchr/testify/blob/master/LICENSE

The following modules are pulled in transitively by the test dependencies above:

- **`github.com/davecgh/go-spew@v1.1.1`** — ISC License — Copyright (c) 2012-2016
  Dave Collins (https://github.com/davecgh/go-spew)
- **`github.com/pmezard/go-difflib@v1.0.0`** — BSD 3-Clause License — Copyright
  (c) 2013 Patrick Mezard (https://github.com/pmezard/go-difflib)
- **`gopkg.in/yaml.v3@v3.0.1`** — Apache License 2.0 and MIT License — Copyright
  (c) 2006-2011 Kirill Simonov and Canonical Ltd. (https://github.com/go-yaml/yaml)

---

## Attribution

If you redistribute this software, in source or binary form, you must retain
the copyright notice and license text of the Craftgate Go Client, together
with the attributions for the third-party components listed above, in
accordance with their respective license terms.

For questions regarding licensing or attribution, contact info@craftgate.io.
