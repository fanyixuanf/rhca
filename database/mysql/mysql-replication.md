## mysql

- 创建mysql基础镜像🚀️ 

  - ```yaml
    mysql-master: &mysql
    container_name: mysql-master
    image: mysql:8.0.21
    restart: always
    ports:
      - "33060:3306"
    environment:
      MYSQL_DATABASE: 'test'
      MYSQL_ROOT_PASSWORD: 'abcd@1234'
    command: [
      '--character-set-server=utf8mb4',
      '--collation-server=utf8mb4_general_ci',
      '--log-bin=mysql-bin',
      '--server-id=1',
      '--sync_binlog=1',
      '--default-authentication-plugin=mysql_native_password',
    ]
    volumes:
      - mysql-master-data:/var/lib/mysql
      - './config/custom.cnf:/etc/mysql/conf.d/custom.cnf'
    networks:
      - mysql-network
    ```
- 创建node镜像🚀️ 

  - ```yaml
    mysql-node-1: &mysql-node
    <<: *mysql
    container_name: mysql-node-1
    environment:
      MYSQL_DATABASE: 'test'
      MYSQL_ROOT_PASSWORD: 'abcd@1234'
    ports:
      - "33061:3306"
    restart: always
    depends_on:
      - mysql-master
    volumes:
      - mysql-node-1-data:/var/lib/mysql
      - './config/custom.cnf:/etc/mysql/conf.d/custom.cnf'
    command: [
      '--character-set-server=utf8mb4',
      '--collation-server=utf8mb4_general_ci',
      '--server-id=10',
      '--default-authentication-plugin=mysql_native_password',
    ]
    networks:
      - mysql-network
    ```
    
- 修改密码策略
  - `ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'abcd@1234';`
  - `FLUSH PRIVILEGES;`

- master创建replication用户
  - `CREATE USER 'replication'@'%' IDENTIFIED BY 'abcd@1234';`
  - `GRANT REPLICATION SLAVE ON database.table TO 'replication'@'%';`
  - `show master status;`
  - `CREATE USER 'replication'@'%' IDENTIFIED WITH 'mysql_native_password' BY 'abcd@1234';`
  - `GRANT REPLICATION SLAVE ON *.* TO 'replication'@'%'; `
  - [2061：Authentication plugin 'caching_sha2_password' reported error:Authentication require secure connection 解决方案](https://www.modb.pro/db/29919)

- node配置
    - 查看容器ip，更新到下面语句中:
    - `change master to master_host='172.16.0.10',master_user='replication',master_password='abcd@1234',master_port=3306,master_log_file='mysql-bin.000003', master_log_pos=832,master_connect_retry=30;`
    - `start slave;`
    - `show slave status;`

## docker
- `docker network inspect mysql_mysql-network` 查看容器IP
  - ```json
    [
    {
        "Name": "mysql_mysql-network",
        "Id": "9af9bf65030ef08051fa279bd00354fcd98dd2f459ff6ed2cd43fb28040a3e98",
        "Created": "2023-02-27T06:44:07.8689055Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.27.0.0/16",
                    "Gateway": "172.27.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": true,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "70bc8c75173491a5129de53bc69c261e1013df9609b7a16c2e37257e35c0df22": {
                "Name": "mysql-node-1",
                "EndpointID": "79ecb6ef37e47580602dd8429409938880288de292a534d8f6e46a1f50826c41",
                "MacAddress": "02:42:ac:1b:00:04",
                "IPv4Address": "172.27.0.4/16",
                "IPv6Address": ""
            },
            "aca69e687ed4099f075a91a465c6cc6a899e169675013d82cf630fc19a9139d1": {
                "Name": "mysql-node-2",
                "EndpointID": "74a29f1b3d7742f6843834ce80e819de198de0a1ae7a2ad37cf2f47a2ddeb2fc",
                "MacAddress": "02:42:ac:1b:00:03",
                "IPv4Address": "172.27.0.3/16",
                "IPv6Address": ""
            },
            "f30730b1fb50c75f8e074dd149a7bd04e486e7af6a4dd7c8b330c7cc6cef7eca": {
                "Name": "mysql-master",
                "EndpointID": "0b916cddece5adf430c663f6df5d47095ce87a802aecc9909fc10b3afecb2f50",
                "MacAddress": "02:42:ac:1b:00:02",
                "IPv4Address": "172.27.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {
            "com.docker.compose.network": "mysql-network",
            "com.docker.compose.project": "mysql",
            "com.docker.compose.version": "1.29.2"
        }
    }
    ]
    ```

## Mysql Router
- docker 安装 `docker pull mysql/mysql-router:8.0.32`
  - `docker-compose.yml`
    - ```yaml
      version: '2'
      services:
          mysql-router:
              environment:
                  MYSQL_HOST: "127.0.0.1"
                  MYSQL_PORT: 33060
                  MYSQL_USER: 'root'
                  MYSQL_PASSWORD: 'abcd@1234'
                  MYSQL_INNODB_CLUSTER_MEMBERS: 3
                  MYSQL_CREATE_ROUTER_USER: 0
              image: mysql/mysql-router:8.0.32
              volumes:
                  - "./mysqlrouter-d.conf:/tmp/myrouter/mysqlrouter.conf"
              command: /bin/bash -c "mysqlrouter --config /tmp/myrouter/mysqlrouter.conf"
              ports:
                - "7003:7003"
                - "7004:7004"
      ```
- win安装：
  - 下载安装包，bin目录加入环境变量
  - 创建配置文件
    - ```
          [DEFAULT]
          config_folder = ./log
          logging_folder = ./log
          runtime_folder = ./log

          [logger]
          level = INFO

          [routing:slaves]
          bind_address = 127.0.0.1:7001
          destinations = 127.0.0.1:33061,127.0.0.1:33062
          mode = read-only
          connect_timeout = 1

          [routing:masters]
          bind_address = 127.0.0.1:7002
          destinations = 127.0.0.1:33060
          mode = read-write
          connect_timeout = 2
      ```
  - 运行`mysqlrouter --config=mysqlrouter.conf`

### MGR(MySQL Group Replication)
- [MYSQL MGR配置](https://blog.csdn.net/wzc900810/article/details/108734130)
- https://www.modb.pro/db/138792
- 构建镜像
- 创建复制账号
  - SET SQL_LOG_BIN=0;
  - CREATE USER rpl_user@'%' IDENTIFIED WITH 'mysql_native_password' BY 'abcd@1234';
  - GRANT REPLICATION SLAVE ON *.* TO 'rpl_user'@'%';
  - FLUSH PRIVILEGES;
  - SET SQL_LOG_BIN=1;
  - CHANGE MASTER TO MASTER_USER='rpl_user', MASTER_PASSWORD='abcd@1234' FOR CHANNEL 'group_replication_recovery';

- 主库执行
  - SET global group_replication_bootstrap_group=ON;
  - start GROUP_REPLICATION;
  - SET global group_replication_bootstrap_group=OFF;
- 从库执行
  - start GROUP_REPLICATION; 
  - SELECT * FROM performance_schema.replication_group_members;
