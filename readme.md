## Demo Gorilla csrf + restful js client

Use `gorilla/mux` and `gorilla/csrf`

### Purpose

Official demos are included in its readme and unit test, which is not fully coded out.

We need to play with chrome network panel to see:

    Cookie
    Headers

It serves as example to further security analysis

### Feature:
1. js solution with `axios` similar to react
2. middleware latest stack demo

### Credits

inspired by [blog](https://qiita.com/obr-note/items/898f862a6ebe27c45de4)

* my innovation:

    - slight mod to display error.

    - add bash

Links:

[owasp](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#use-of-custom-request-headers)

[react csrf](https://www.stackhawk.com/blog/react-csrf-protection-guide-examples-and-how-to-enable-it/)

Stateless:

[poc java - done by owasp volunteer](https://github.com/righettod/poc-csrf)

Stateful:

[OWASP-CSRFGuard](https://github.com/aramrami/OWASP-CSRFGuard)


