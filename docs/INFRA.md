1. AWS VPC (aws_vpc)

    **Purpose:** The VPC (Virtual Private Cloud) provides a logically isolated section of the AWS cloud where you can launch AWS resources.

2. Internet Gateway (aws_internet_gateway)

    **Purpose:** The Internet Gateway enables communication between instances in your VPC and the internet.

3. Route Table (aws_route_table)

    **Purpose:** A route table contains a set of rules, called routes, that are used to determine where network traffic from your subnet or gateway is directed.

4. Route Table Associations

    **Purpose:** These resources associate the route table with the defined subnets.

5. Subnets (aws_subnet)

    **Purpose:** Subnets are segments of your VPC's IP address range where you can place groups of isolated resources.

6. Security Group (aws_security_group)

    **Purpose:** Security groups act as virtual firewalls that control inbound and outbound traffic to AWS resources.

7. EC2 Key Pair (aws_key_pair)

    **Purpose:** A key pair allows you to securely connect to your EC2 instances.

8. Application Load Balancer (aws_lb)

    **Purpose:** The ALB distributes incoming application traffic across multiple targets (like EC2 instances).

9. EC2 Instance (aws_instance)

    **Purpose:** EC2 instances are virtual servers in the cloud that you can use to run applications.
