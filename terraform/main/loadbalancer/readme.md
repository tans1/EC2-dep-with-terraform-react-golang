here we have
1. alb configuration
2. target groups (both react and golang)
3. listener for both port 80 and 443
5. listener rules to route based on the header 
        - example.com to react
        - api.example.com to golang


notes
listener's priority :- When an Application Load Balancer (ALB) receives a request, it needs to determine which listener rule to apply. The priority attribute specifies the order in which the rules are evaluated. Rules with lower priority values are evaluated before rules with higher priority values.