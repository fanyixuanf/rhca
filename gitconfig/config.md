# Configure multiple Git accounts on the same computer

> 为同一个电脑，配置多个 git 账号，其整体流程如下：
> 
> 清空默认的全局 user.name 和 user.email
> 
> 为不同的 git 账户生成不同的 ssh-key
> 
> 将以上的 ssh-key 分别添加到 ssh-agent 信任列表
> 
> 添加以上的公钥到自己的 git 账户中
> 
> 在 config 文件配置多个 ssh-key

### 1. 清空默认的全局 user.name 和 user.email

```bash
git config --global --unset user.name
git config --global --unset user.email
```

### 2. 为不同的 git 账户生成不同的 ssh-key
`id_ras` 是默认的文件名称，我们现在需要生成不同的 `ssh-key`，所以要设置不同的文件存储对应的公钥，比如：自己的 GitHub 账户，使用 `id_ras_github` 命名；公司的账户，使用 `id_ras_company` 来命名

```bash
ssh-keygen -t id_ras_github -C "xxx@xx.com"
ssh-keygen -t id_ras_company -C "xxx@company.com"
```

### 3. 将 ssh-key 分别添加到 ssh-agent 信任列表
```bash
ssh-add ~/.ssh/id_ras_github
ssh-add ~/.ssh/id_ras_company
```
> 如果看到 Identitiy added: ~/.ssh/id_ras_github，就表示添加成功了

### 4. 添加公钥到自己的 git 账户中
> 使用命令，copy公钥，到 git 账户中粘贴即可
```bash
pbcopy < ~/.ssh/id_ras_github.pub
pbcopy < ~/.ssh/id_ras_company.pub
```

### 5. 在 config 文件配置多个 ssh-key
> 在 .ssh/ 目录下，config文件（没有的话新建一个）
```bash
# github
Host github.com
    HostName github.com
    PreferredAuthentications publickey
    IdentityFile C:\\Users\\\fanyixuan\\.ssh\\id_ras_github
# github
Host 192.168.x.x
    HostName 192.168.x.x
    PreferredAuthentications publickey
    IdentityFile C:\\Users\\\fanyixuan\\.ssh\\x

# github 可添加代理
Host 192.168.x.x
    HostName 192.168.x.x
    PreferredAuthentications publickey
    IdentityFile C:\\Users\\\fanyixuan\\.ssh\\id_ras_company
    ProxyCommand connect -H proxy.com:port %h %p
    ForwardAgent yes
```

### 最后
> 在不同的代码仓库进行代码提交时，记得检查用户名和邮箱，以免混淆。设置用户名和邮箱的命令如下：
```bash
git config --local user.name xxx
git config --local user.email xxx@xxx.com
```