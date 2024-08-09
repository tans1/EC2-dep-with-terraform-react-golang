resource "aws_lb" "loadbalancer" {
  name               = "ec2-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [var.security_group_id]
  subnets            = var.public_subnets

  tags = {
    Name = "ec2-alb"
  }
}

################## TARGET GROUPS #########################
resource "aws_lb_target_group" "react_target_group" {
  name        = "react-target-group"
  port        = 3000
  protocol    = "HTTP"
  target_type = "instance"
  vpc_id      = var.vpc_id

  health_check {
    path = "/"
  }
}

resource "aws_lb_target_group" "golang_target_group" {
  name        = "golang-target-group"
  port        = 4000
  protocol    = "HTTP"
  target_type = "instance"
  vpc_id      = var.vpc_id

  health_check {
    path = "/"
  }
}

################ LISTENERS ##############################
resource "aws_lb_listener" "alb_http_listener" {
  load_balancer_arn = aws_lb.loadbalancer.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"
    redirect {
      protocol    = "HTTPS"
      port        = "443"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_listener" "alb_https_listener" {
  load_balancer_arn = aws_lb.loadbalancer.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = var.ssl_policy
  certificate_arn   = var.ssl_certificate_arn

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "404: Not Found"
      status_code  = "404"
    }

  }
}

################## LISTNER RULES ######################################
resource "aws_lb_listener_rule" "react_https" {
  listener_arn = aws_lb_listener.alb_https_listener.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.react_target_group.arn
  }

  condition {
    host_header {
      values = [var.main_domain, "www.${var.main_domain}"]
    }
  }
}

resource "aws_lb_listener_rule" "golang_https" {
  listener_arn = aws_lb_listener.alb_https_listener.arn
  priority     = 200

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.golang_target_group.arn
  }

  condition {
    host_header {
      values = [var.api_domain]
    }
  }
}


################## Loadbalancer attachment to instance ######################
resource "aws_lb_target_group_attachment" "react" {
  target_group_arn = aws_lb_target_group.react_target_group.arn
  target_id        = var.instance_id
  port             = 3000
}

resource "aws_lb_target_group_attachment" "golang" {
  target_group_arn = aws_lb_target_group.golang_target_group.arn
  target_id        = var.instance_id
  port             = 4000
}