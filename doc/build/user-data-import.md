# 在fleetmanager数据库中增加信息

## 数据说明：
   domain_id: 账号对应的账号id
   domain_name: 账号对应的账号ming
   IAM_id: 账号对应的IAM用户id
   IAM_name: 账号对应的IAM用户ID
## 数据导入过程：
1. 向`res_domain`中插入委托账号(即账号2)的信息：
   ```sql
   INSERT INTO res_domain (
      id, 
      origin_domain_name, 
      origin_domain_id, 
      res_domain_name,
      res_domain_id, 
      creation_time, 
      region
   )
   VALUES (
      '1', 
      '资源账号的帐号名', 
      '资源账号的帐号ID', 
      '资源账号的帐号名', 
      '资源账号的帐号ID', 
      '2022-09-21 16:10:00', 
      'region'
   )
   ```
2. 向`res_user`中插入委托账号的用户信息：
   ```sql
   INSERT INTO res_user (
      id, 
      origin_domain_id, 
      res_domain_id, 
      res_user_name,
      res_user_id, 
      res_user_pw, 
      creation_time, 
      update_time, 
      region
   )
   VALUES (
      '1', 
      '资源账号的帐号ID', 
      '资源账号的帐号ID', 
      '资源账号的IAM用户名', 
      '资源账号的IAM用户ID', 
      '', 
      '2022-09-21 16:10:00', 
      '2022-09-21 16:10:00', 
      'region'
   )
   ```
3. 向`res_project`中插入委托账号的项目信息：

   ```sql
   INSERT INTO res_project (
      id, 
      origin_domain_id, 
      origin_project_id, 
      res_domain_id, 
      res_project_id, 
      creation_time, 
      region
   )
   VALUES (
      '1', 
      '资源账号的帐号ID', 
      '资源账号在指定region的项目ID', 
      '资源账号的帐号ID', 
      '资源账号在指定region的项目ID', 
      '2022-09-21 16:10:00', 
      'region')
   ```
4. 向`res_keypair`中插入委托账号2的密匙对信息：
   ```sql
   INSERT INTO res_keypair (
      id, 
      origin_domain_id, 
      res_domain_id, 
      keypair_name, 
      creation_time, 
      update_time, 
      region
   )
   VALUES (
      '1', 
      '资源账号的帐号ID', 
      '资源账号的帐号ID', 
      '资源账号的账号密钥对名称', 
      '2022-09-21 16:10:00', 
      '2022-09-21 16:10:00', 
      'region')
   ```
   
5. 在`res_agency`委托账号对对被委托账号的委托信息：
   ```sql
   INSERT INTO res_agency (
      id, 
      origin_domain_id, 
      res_domain_id, 
      res_user_id, 
      agency_name,
      creation_time, 
      region)
   VALUES (
      '1', 
      '资源账号的帐号ID', 
      '资源账号的帐号ID', 
      '资源账号的IAM用户ID', 
      '资源账号对账号1的委托名称', 
      '2022-09-21 16:10:00', 
      'region'
   )
   ```
6. 在`build`中插入应用构建的信息:
   ```sql
   INSERT INTO build (
      id, 
      project_id, 
      state, 
      image_id, 
      image_region, 
      creation_time, 
      update_time, 
      termination_time
   )
   VALUES (
      '1', 
      '资源账号在指定region的项目ID', 
      'READY', 
      '资源账号的应用镜像ID', 
      'region', 
      '2022-09-21 16:10:00', 
      '2022-09-21 16:10:00', 
      '0000-00-00 00:00:00'
   )
   ```
7. 在build_image中插入镜像信息: 
   ```sql
   INSERT INTO build_image (
      id, 
      project_id, 
      build_id, 
      image_id, 
      image_region_id, 
      create_time
   )
   VALUES (
      '1', 
      '资源账号在指定region下的项目ID', 
      'build_id', 
      '资源账号的应用镜像ID', 
      'region', 
      '2022-09-21 16:10:00'
   )
   ```