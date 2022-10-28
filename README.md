# aws product learn

## EC2
    ssh -i C:\Users\Administrator\OneDrive\study\aws-product-learn\aws-product-learn.pem ubuntu@16.163.58.102

## S3
    最大单个文件5TB
    持久性11个9
    S3 Glacier不需要立即访问但需要灵活

## VPC
    VPC可以跨az
    VPC不可以跨regions
    S3不属于VPC内部

  - 典型VPC网络架构

    ![典型VPC网络架构](./vpcArchitecture/vpc-architecture_diagram.png)

  - 高可用VPC网络架构

    ![高可用VPC网络架构](./vpcHAArchitecture/vpc-architecture_diagram.png)

## CloudWatch
    默认监控指标
      CPU
      IO
      Network
      Memery默认不监控，要手动开户

## ELB Elastic Load Balancing
    ELB
      高可用性
      可以检测运行状态
      SSL/TLS终止
      运行监控
    2代三类产品
      ALB Application Load Balancing
        LV7
      NLB Network Load Balancing
        LV4
        100W级并发
        2019年支持udp
        
## RDS Relational Database Service
    支持主流数据库
      Amazon Aurora
      Postgresql
      MariaDB
      Oracle
      MySql
      SqlServer
    特点
      轻松扩展
      自动软件修复
      自动备份 默认开户自动备份策略
      数据库快照 基于EBS块存储snapshop
      多regions AZ部署
      自动主备更换 master replica(slave歧视)
      静态和传输中加密 数据，log加密

## DynamoDB适用场景
    无服务器Web应用程序
    微服务数据存储
    移动后端
    广告技术
    游戏
    物联网Iot

## Redshift OLAP On-Line Analytical Processing
    机器学习
    大规模并行查询
    列式存储
    BI数据分析决策

## CloudFormation
    IAC Infrastructure As Code
    云基础设施建模并对其进行预置

## Elastic Beanstalk代码即服务
    简单快速
    资源预置
    负载均衡
    自动扩展
    监控，日志记录和跟踪
    各种应用程序部署选项

## Diretc Connect
    专线
  ![专线网络架构](./directConnect/1666883898115.jpg)
