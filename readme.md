## Demo REST backend Gorilla csrf middleware and Js frontend

Use `gorilla/mux` and `gorilla/csrf`

### Purpose

Official demos are included in its readme and unit test, which is not fully coded out.

For user who want to see `chrome network panel` breakdowns in parts,

    Cookie
    Headers
    Custome Header

- It serves as example to further security analysis

  You might be interested in fuzz test the encryption :)

### Major feature

1. Demoed backend Restful api.
2. Demoed frontend js for further react adaption.
   - It helps none-mvc, react, mobile developers to quickly learn csrf practice.
3. Demoed middleware in brunch `diy_middleware`, 
   - so backend ppl can make their own middleware.
4. How to unit test this stuff. Actually, very good place to start learning concept.
   - it requires minimal effort to run and proof correctness.
   
### Credits

Inspired by this japanese [blog](https://qiita.com/obr-note/items/898f862a6ebe27c45de4)

### My contrib on top of the blog result

    - html input take dynamic name. 
      - Let you feel the change rather than hard-value.
    - more server error hanlding, to know where not work.
    - add bash to walk pass error on start.
    - use several branch to demo the use of csrf with popular framework `mux` 
      or just with `http`. Showed a more flexible and orthogonal solution.

### Reference

[owasp](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html#use-of-custom-request-headers)

[react csrf](https://www.stackhawk.com/blog/react-csrf-protection-guide-examples-and-how-to-enable-it/)

Stateless:

[poc java - done by owasp volunteer](https://github.com/righettod/poc-csrf)

Stateful:

[OWASP-CSRFGuard](https://github.com/aramrami/OWASP-CSRFGuard)


