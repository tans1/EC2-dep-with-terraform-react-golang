## Note
here I am deploying the project on both EC2


I used
golang(gin and mysql)- backend
redis - catching
react(vite, shadcn and tailwind) - frontend
alb - webserver and loadbalancer


Speaking of this, If we used different repositories for each of the following service, it is also better to use another repository for the terraform to make the code more clear and maintainable.

=> to login to the instance I am using SSM , because
    1. since the instance is in the private subnet, I cannot use SSH
    2. bastion comes with risks, like :- single point of failure, security vulnerability ...
    
=> and also, I am using ubuntu Os with the t3.micro AMI.