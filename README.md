![image](https://github.com/user-attachments/assets/9dbaf763-4989-4190-805a-11afe63dc284)

> This is a shadow testing utility that'll help developers perform shadow testing for their re-writes :)

**Here's how it'll work**

![image](https://github.com/user-attachments/assets/09c93e87-29f7-4838-8dda-561bab5f8973)


- The nginx server will have reverse proxies for each service
- The HTTP request will be intercepted by a go wasm filter
- using a yaml configured port mapping, it'll be redirected to multiple services instead of just once
- the responses will be stored/sent to prometheous
- grafana dashboards to read said data using promQL

note: whether to use nginx or keep it simple with a straightforward single binary approach is TBD. 
