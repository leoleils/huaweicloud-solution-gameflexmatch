# Content
- [Content](#content)
- [1. Introduction to GameFlexMatch](#1-introduction-to-gameflexmatch)
  - [1.1. Component Function](#11-component-function)
    - [1.1.1. The FleetManager component provides the following functions:](#111-the-fleetmanager-component-provides-the-following-functions)
    - [1.1.2. Functions of the AASS component:](#112-functions-of-the-aass-component)
    - [1.1.3. AppGateway component functions:](#113-appgateway-component-functions)
    - [1.1.4. Functions of the AuxProxy component](#114-functions-of-the-auxproxy-component)
  - [1.2. Console page](#12-console-page)
- [2. Home Page](#2-home-page)
  - [2.1. Brief Introduction](#21-brief-introduction)
  - [2.2. Operation](#22-operation)
- [3. User management](#3-user-management)
  - [3.1. Initial Login Operation](#31-initial-login-operation)
  - [3.2. User management](#32-user-management)
    - [3.2.1. Creating a User](#321-creating-a-user)
    - [3.2.2. Modify a User](#322-modify-a-user)
    - [3.2.3. Delete user](#323-delete-user)
  - [3.3. Tenant management](#33-tenant-management)
    - [3.3.1. Creating a Tenant](#331-creating-a-tenant)
    - [3.3.2. Querying a Tenant](#332-querying-a-tenant)
    - [3.3.3. Binding a Tenant](#333-binding-a-tenant)
    - [3.3.4. Switch a Tenant](#334-switch-a-tenant)
    - [3.3.5. Delete a Tenant](#335-delete-a-tenant)
- [4. Application package management](#4-application-package-management)
  - [4.1. Creating an Application Package](#41-creating-an-application-package)
    - [4.1.1. Operation Scenario](#411-operation-scenario)
    - [4.1.2. Creation Instructions](#412-creation-instructions)
    - [4.1.3. Operation Procedure](#413-operation-procedure)
  - [4.2. Modifying an Application Package](#42-modifying-an-application-package)
    - [4.2.1. Operation Scenario](#421-operation-scenario)
    - [4.2.2. Operation Procedure](#422-operation-procedure)
  - [4.3. Delete an application package](#43-delete-an-application-package)
    - [4.3.1. Operation Scenario](#431-operation-scenario)
    - [4.3.2. Delete Instructions](#432-delete-instructions)
    - [4.3.3. Operation Procedure](#433-operation-procedure)
- [5. Creating a DB Instance Flavor Group](#5-creating-a-db-instance-flavor-group)
  - [5.1. Creating a VM Instance Specification Group](#51-creating-a-vm-instance-specification-group)
    - [5.1.1. Operation Scenario](#511-operation-scenario)
    - [5.1.2. Operation Procedure](#512-operation-procedure)
  - [5.2. Creating a Pod Instance Flavor Group](#52-creating-a-pod-instance-flavor-group)
    - [5.2.1. Operation Scenario](#521-operation-scenario)
    - [5.2.2. Operation Procedure](#522-operation-procedure)
- [6. Process Queue Management](#6-process-queue-management)
  - [6.1. Create Apply Process Queue](#61-create-apply-process-queue)
    - [6.1.1. Operation Scenario](#611-operation-scenario)
    - [6.1.2. Operation Procedure](#612-operation-procedure)
  - [6.2. Modifying the application process queue](#62-modifying-the-application-process-queue)
    - [6.2.1. Operation Scenario](#621-operation-scenario)
    - [6.2.2. Operation Procedure](#622-operation-procedure)
  - [6.3. Deleting the Apply Process Queue](#63-deleting-the-apply-process-queue)
    - [6.3.1. Operation Scenario](#631-operation-scenario)
    - [6.3.2. Delete Instructions](#632-delete-instructions)
    - [6.3.3. Operation Procedure](#633-operation-procedure)
- [7. Auto-Scale policy management](#7-auto-scale-policy-management)
  - [7.1. Creating an Scaling Policy](#71-creating-an-scaling-policy)
    - [7.1.1. Operation Scenario](#711-operation-scenario)
    - [7.1.2. Operation Procedure](#712-operation-procedure)
  - [7.2. Modifying an Scaling Policy](#72-modifying-an-scaling-policy)
    - [7.2.1. Operation Scenario](#721-operation-scenario)
    - [7.2.2. Operation Procedure](#722-operation-procedure)
  - [7.3. Deleting an Scaling Policy](#73-deleting-an-scaling-policy)
    - [7.3.1. Operation Scenario](#731-operation-scenario)
    - [7.3.2. Operation Procedure](#732-operation-procedure)
- [8. Alias Management](#8-alias-management)
  - [8.1. Creating aliases](#81-creating-aliases)
    - [8.1.1. Operation Scenario](#811-operation-scenario)
    - [8.1.2. Creation Instructions](#812-creation-instructions)
    - [8.1.3. Operation Procedure](#813-operation-procedure)
  - [8.2. Modifying aliases](#82-modifying-aliases)
    - [8.2.1. Operation Scenario](#821-operation-scenario)
    - [8.2.2. Operation Procedure](#822-operation-procedure)
  - [8.3. Querying aliases](#83-querying-aliases)
    - [8.3.1. Operation Scenario](#831-operation-scenario)
    - [8.3.2. Operation Procedure](#832-operation-procedure)
  - [8.4. Delete Alias](#84-delete-alias)
    - [8.4.1. Operation Scenario](#841-operation-scenario)
    - [8.4.2. Operation Procedure](#842-operation-procedure)
- [9. Log management](#9-log-management)
  - [9.1. Creating Log Access](#91-creating-log-access)
    - [9.1.1. Operation Scenario](#911-operation-scenario)
    - [9.1.2. Operation Procedure](#912-operation-procedure)
  - [9.2. Creating a Log Dump](#92-creating-a-log-dump)
    - [9.2.1. Operation Scenario](#921-operation-scenario)
    - [9.2.2. Operation Procedure](#922-operation-procedure)
  - [9.3. Deleting Log Access and Dump](#93-deleting-log-access-and-dump)
    - [9.3.1. Background](#931-background)
    - [9.3.2. Operation Procedure](#932-operation-procedure)
- [10. Event Audit](#10-event-audit)
  - [10.1. Viewing the Event List](#101-viewing-the-event-list)
    - [10.1.1. Operation Scenarios](#1011-operation-scenarios)
    - [10.1.2. Operation Procedure](#1012-operation-procedure)
  - [10.2. 11.2 Message Management](#102-112-message-management)
    - [10.2.1. 11.2.1 Operation Scenarios](#1021-1121-operation-scenarios)
    - [10.2.2. 11.2.2 Operation Procedure](#1022-1122-operation-procedure)
- [11. API invoking](#11-api-invoking)
  - [11.1. Interface authentication](#111-interface-authentication)
- [12. FAQs](#12-faqs)
  - [12.1. Fleet and Application Package Creation](#121-fleet-and-application-package-creation)
  - [12.2. Viewing Session Error Causes](#122-viewing-session-error-causes)
  - [12.3. Version iteration](#123-version-iteration)
    - [12.3.1. Step by step traffic diversion](#1231-step-by-step-traffic-diversion)
    - [12.3.2. Full traffic diversion](#1232-full-traffic-diversion)
    - [12.3.3. Other questions](#1233-other-questions)


# 1. Introduction to GameFlexMatch

This service consists of four components: FleetManager, AppGateway,
AASS, and AuxProxy. FleetManager is a global component, AppGateway and
AASS are region-level components, and AuxProxy is a VM-level component.
The following figure shows the overall architecture.

![](images-en/media/image1.png)

## 1.1. Component Function

### 1.1.1. The FleetManager component provides the following functions:

Deploys and manages applications globally and dynamically, and supports
the configuration of dynamic deployment policies.

### 1.1.2. Functions of the AASS component:

Creates applications and manages and executes scaling policies.

-   Determine whether to perform application process resiliency based on
    the system load.

-   Determine the capacity expansion/reduction process and resource
    quantity based on the number of service requests.

-   Interworks with a third-party resource management system to control
    resource application and release.

### 1.1.3. AppGateway component functions:

-   Stores and manages sessions (server sessions and client sessions)
    objects (creation, deletion, change, and status statistics), and
    collects statistics on client request status.

-   Manages regional processes (mainly statistics and display).

-   Reporting statistics (including statistics on processes, client
    sessions, server sessions, and client requests)

-   Affinity rule management (Affinity rules are the mapping between the
    client request and the server session, and are used to select the
    application process object of the client request for the
    gateway-proxy.)

### 1.1.4. Functions of the AuxProxy component

Management plane component of the host. One component is started on each
VM and manages application processes on the VM.

-   Manages the life cycle of application processes.

-   Monitors application processes and reports the status information
    about application processes, client sessions, and server sessions.

## 1.2. Console page

This project provides frontend pages developed based on the Vue to
facilitate users to operate and maintain service components and cloud
resources.



# 2. Home Page

## 2.1. Brief Introduction

-   The home page displays the overall running status of the
    GameFlexMatch tenant in a region, including the number of Fleets,
    instances, processes, and sessions.

-   Displays the number of instances, processes, and sessions in each
    Fleet in a chart.

-   Displays the number of instances, processes, and sessions in
    different states in each Fleet in a chart.

## 2.2. Operation

· Click the homepage to view the overall running status.

![](images-en/media/image29.png)

· To view the running status of several Fleets, choose Home \> Fleet
Running Status. You can select multiple Fleets.

# 3. User management

GameFlexMatch uses GameFlexMatch users to manage HUAWEI CLOUD tenants.
Users can be associated with tenants to manage HUAWEI CLOUD resources.
Multiple users can be associated with one tenant, and a user can be
associated with multiple tenants. Instant switchover is supported.

![](images-en/media/image30.png)

## 3.1. Initial Login Operation

1.  When the project is started, the system checks whether the database
    has an account. If no account exists, the system initializes a super
    administrator account. The default password is the value of
    "DEFAULT_LOGIN_PASSWORD" in the fleetmanager_run.sh script.

2.  On the login page, enter the user name "admin" and password.

3.  You need to change the default password upon the first login. If the
    password is changed successfully, the system automatically logs out
    and logs in again.

4.  After the login is successful, the response data contains
    Auth-Token. All APIs except the login API must carry Auth-Token in
    the request header for authentication (see chapter 12 API invoking).

5.  Description:

    a.  For the token generated based on the JWT architecture, you can
        set the key for generating the token and the expiration time in
        the environment variable. The corresponding fields are JWTKEY
        and JWT_TOKEN_LIFETIME.

    b.  A login operation is a session request. The session expiration
        time can be set by the LOGIN_SESSION_LIFETIME field in the
        script file. If the session expires, the session will not be
        renewed and you need to log in again.

6.  Frozen accounts cannot be used to log in to the system.

7.  After you log out, the current cookies will be deleted and the
    original token cannot be used.

Description of user types

| **User  type**      | **Value** |
| ------------------- | --------- |
| Read-only user      | 0         |
| Common user         | 5         |
| Super administrator | 9         |

Description of the activation state

| **Activated  state**                                         | **Value** |
| ------------------------------------------------------------ | --------- |
| Inactive. The password is not changed  upon the first login. | 0         |
| Activated, password changed                                  | 1         |
| If the account is frozen and the password  is incorrect for five consecutive times, the administrator needs to reset the  password. | - 1       |

## 3.2. User management

### 3.2.1. Creating a User

Administrators can create sub-users to manage the platform. Common users
and read-only users are supported. Common users have the permission to
modify resources on the platform. Read-only users can only view
resources on the platform.

-   Administrator Creating a User.

1.  Administrator Login

2.  Find the entry for adding a user.

> ![](images-en/media/image31.png)

3.  Add a common user. NewUser is used as an example.

> ![](images-en/media/image32.png)

4.  The administrator associates the resource tenant with the new user.

> ![](images-en/media/image33.png)

5.  Viewing user Details

> ![](images-en/media/image34.png)

6.  Viewing Resource Tenant Information Associated with current user.

![](images-en/media/image35.png)

### 3.2.2. Modify a User

A user can modify only personal information (email address and phone
number). The super administrator can modify the personal information,
activation status, and user type of a user, and reset the password.

1\. Log in to the console as the administrator.

2\. Choose Configuration \> User Management, select the user to be
modified, and click Modify or Reset Password.

![](images-en/media/image36.png)

### 3.2.3. Delete user

1.  Log in to the console as the administrator.

2.  Choose Configuration \> User Management, select the user to be
    deleted, and click Delete.

![](images-en/media/image37.png)

## 3.3. Tenant management

Tenant management is an important part of GameFlexMatch. Platform users
create and manage HUAWEI CLOUD resources through associated HUAWEI CLOUD
tenants.

### 3.3.1. Creating a Tenant

You can create a tenant in any of the following modes:

-   The administrator logs in to the system for the first time.

-   The administrator adds a user.

-   The administrator adds a tenant for the user on the console.

To add a tenant, perform the following steps:

1.  Choose Configuration \> User Management and add a tenant for the
    target user.

![](images-en/media/image38.png)

2.  Parameters for creating a tenant are as follows:

Use the tenant information:

-   Tenant name: name of the GameFlexMatch platform tenant information

-   Project ID: specifies the ID of the API project used by
    GameFlexMatch to create resources. You can obtain the project ID by
    choosing My Credential \> API Credential on the HUAWEI CLOUD
    console.

-   Access Key: specifies the key used by GameFlexMatch to access APIs.
    You can obtain the key by choosing My Credential \> Access Key on
    the HUAWEI CLOUD console.

-   Secret Access Key: specifies the key used by GameFlexMatch to access
    APIs. You can obtain the key by choosing My Credential \> Access Key
    on the HUAWEI CLOUD console.

-   Region: indicates the region where the resource tenant creates
    GameFlexMatch resources.

-   Key name: Login authentication key used during VM scaling by
    GameFlexMatch

-   Cloud Service Agency Name: This parameter is used to install ICAgent
    during image packing and use LTS to dump logs.

Application package information (Optional. If this parameter is left
blank, the default parameter is used.)

-   Auxproxy Path: OBS path for storing Auxproxy files.

-   ECS-Application Package Configuration Script Path: OBS path for
    storing the script file used to create ECS resource application
    packages.

-   Region: indicates the region where the OBS bucket is located.

-   Container-Application Package Configuration Script Path: OBS path
    for storing the script for creating container resource application
    packages. If a container is used, this parameter is mandatory.

3.  Click Create to complete the creation process.

### 3.3.2. Querying a Tenant

1.  Log in to the platform as the administrator.

2.  Choose Configuration \> User Management and select a user.

3.  Click the resource tenant details to view the details.

![](images-en/media/image39.png)

### 3.3.3. Binding a Tenant

The administrator binds an existing tenant to the user so that the
tenant information can be reused.

![](images-en/media/image40.png)

### 3.3.4. Switch a Tenant

1.  Logging In to the GameFlexMatch Console

2.  Click in the upper right corner to switch to the associated resource
    tenant.

![](images-en/media/image41.png)

### 3.3.5. Delete a Tenant

1.  Log in to the console as the administrator.

2.  On the User Management page, select a user and select the tenant to
    be deleted in the Account Information area.

![](images-en/media/image42.png)

# 4. Application package management

## 4.1. Creating an Application Package

### 4.1.1. Operation Scenario

You can upload an application package to be hosted to the GameFlexMatch
platform to automatically generate an application package image with the
monitoring and reporting plug-ins installed and use the image to create
a fleet. The VM and POD image packaging modes are supported. You can
also bind an existing private image in HUAWEI CLOUD IMS or SWR to
quickly build a fleet.

### 4.1.2. Creation Instructions

Before creating an application package by uploading a file or using OBS,
create the required VPC and subnet under the current resource tenant.

### 4.1.3. Operation Procedure

1.  Upload the following files to the OBS bucket of the management
    account:

    a.  The Auxproxy service component package, which can be in ZIP or
        RAR format. (Ensure that each file is stored in the same
        directory and does not need to be contained in the auxproxy
        folder.)

    b.  Image packaging environment building scripts "image_env.sh" (VM
        instance image) and "docker_image_env.sh" (Pod instance image)

    c.  the application package (zip or rar) to the OBS bucket of the
        resource account.

2.  In the application image, the application is decompressed in the
    /local/app/{Name of the created application package} directory by
    default. Adjust the application startup script based on the
    directory or modify the application download path in the environment
    building script.

3.  Add the following parameters to the FleetManage_run.sh script for
    starting the FleetManger server components:

> OBS path of the environment build script:
>
> export DEFAULT_AUXPROXY_PATH=Bucket name/auxproxy.zip
>
> OBS Path of the Auxproxy Service Component Package
>
> export DEFAULT_SCRIPT_PATH=Bucket name/image_env.sh

4.  Log in to the GameFlexMatch management console.

5.  Choose Application Package Management \> Create Application Package.

6.  Select the required creation mode and set parameters such as the
    application package name and version number. The following table
    describes the key parameters.

| **Parameters**           | **Explained**                                                | **Example Value** |
| ------------------------ | ------------------------------------------------------------ | ----------------- |
| Application Package Name | Indicates the name of the  created application package. The name contains 2 to 50 characters, which  starts with a letter and consists of letters, digits, underscores (_), and  hyphens (-). | server            |
| Version Number           | Version number of the  created application package. The value contains 1 to 50 characters. | 1.0.0             |
| Application Package Type | Specifies the type of the  created application package. If the value is POD, the SWR image is created.  If the value is VM, the IMS image is created. | VM                |
| VPC                      | Specifies the VPC ID, which  is used to bind the ECS where the image is created. | -                 |
| Subnet                   | Specifies the subnet ID,  which is used to bind the ECS for creating an image. | -                 |
| Operating system         | If Application Package Type  is set to VM, the ECS OS used for creating an image is the underlying OS.  CentOS is recommended. To view the OS, choose HUAWEI CLOUD Console > IMS  > Public Images. When POD is selected for Application Package Type, the  value is the underlying operating system in the container. | CentOS 7.2 64bit  |
| OBS bucket               | OBS bucket name, which is  used to store application package files. | GameFlexMatch     |
| Storage Path             | OBS application package  file path                           | build/server.zip  |

7.  After the parameters are set, click Create.

8.  When the application package status is Ready, it can be used to
    create a fleet.

## 4.2. Modifying an Application Package

### 4.2.1. Operation Scenario

You can modify an application package as required. You can modify the
name, version number, and description of an application package.

### 4.2.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Application Package Management.

3.  In the application package list, locate the row that contains the
    target application package, and click Details.

![](images-en/media/image43.png)

4.  Click Modify next to the parameter and enter the content to be
    modified in the text box.

![](images-en/media/image44.png)

5.  Click OK.

## 4.3. Delete an application package

### 4.3.1. Operation Scenario

You can delete an application package that you no longer need.

### 4.3.2. Delete Instructions

Before deleting the application package, check whether the application
package is associated with a fleet that is not in the terminated state.
If yes, the application package cannot be deleted.

### 4.3.3. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Application Package Management.

3.  In the application package list, locate the row that contains the
    target application package, and click Delete.

> ![](images-en/media/image45.png)

4.  In the displayed dialog box, click OK.

# 5. Creating a DB Instance Flavor Group

The GameFlexMatch platform supports the creation of instance
specification group templates in advance. You can select the template
when creating a fleet.

## 5.1. Creating a VM Instance Specification Group

### 5.1.1. Operation Scenario

For ECS instance specifications with the fleet type as ECS, a maximum of
10 ECS models can be selected under the same flavor. The fleet is
created based on the selected flavor priorities.

### 5.1.2. Operation Procedure

1.  Choose Configuration \> Instance Specification Group and click
    Create Instance Flavor Group.

> ![](images-en/media/image46.png)

2.  Select the VM type instance, enter the group name, select the
    required instance specifications in sequence, and click OK.

![](images-en/media/image47.png)

## 5.2. Creating a Pod Instance Flavor Group

### 5.2.1. Operation Scenario

Allows users to create a specification group of the pod type, select the
VM type, and select the CPU and memory size of the pod instance.

### 5.2.2. Operation Procedure

1.  Choose Configuration \> Instance Specification Group and click
    Create Instance Specification Group.

2.  Select the VM type instance, enter the group name, select the
    required instance specifications in sequence, and click OK.

![](images-en/media/image48.png)

# 6. Process Queue Management

## 6.1. Create Apply Process Queue

### 6.1.1. Operation Scenario

The GameFlexMatch platform provides global dynamic deployment and
management of application processes. You can create application process
queues to carry your service applications and provide services for
external systems. VM or POD instances can be created.

### 6.1.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Fleet Management \> My Fleet \> Create Fleet.

> ![](images-en/media/image49.png)

3.  Configure parameters such as basic information and resource
    information. You can import parameters in a JSON file or export a
    JSON file based on the entered parameters. The following table
    describes the key parameters.

| **Parameters**           | **Explained**                                                | **Example Value** |
| ------------------------ | ------------------------------------------------------------ | ----------------- |
| Application Package Name | Indicates the name of the  created application package. The name contains 2 to 50 characters, which  starts with a letter and consists of letters, digits, underscores (_), and  hyphens (-). | server            |
| Version Number           | Version number of the  created application package. The value contains 1 to 50 characters. | 1.0.0             |
| Application Package Type | Specifies the type of the  created application package. If the value is POD, the SWR image is created.  If the value is VM, the IMS image is created. | VM                |
| VPC                      | Specifies the VPC ID, which  is used to bind the ECS where the image is created. | -                 |
| Subnet                   | Specifies the subnet ID,  which is used to bind the ECS for creating an image. | -                 |
| Operating system         | If Application Package Type  is set to VM, the ECS OS used for creating an image is the underlying OS.  CentOS is recommended. To view the OS, choose HUAWEI CLOUD Console > IMS  > Public Images. When POD is selected for Application Package Type, the  value is the underlying operating system in the container. | CentOS 7.2 64bit  |
| OBS bucket               | OBS bucket name, which is  used to store application package files. | GameFlexMatch     |
| Storage Path             | OBS application package  file path                           | build/server.zip  |

4.  After the parameters are set, click Create.

5.  Wait for the creation. If the Fleet status is Active, the Fleet is
    successfully created. If the Fleet status is Abnormal, locate the
    fault by referring to the fault guide.

## 6.2. Modifying the application process queue

### 6.2.1. Operation Scenario

When managing application process queues, you can modify the basic
information, inbound rules, running configuration, and capacity
information about the application process queues as required.

### 6.2.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Fleet Management \> My Fleet.

3.  In the Fleet list, locate the row that contains the target Fleet,
    and click Details.

> ![](images-en/media/image50.png)

4.  Click "Edit" and enter the content to be modified in the dialog box
    that is displayed.

> ![](images-en/media/image51.png)

5.  Click Save.

## 6.3. Deleting the Apply Process Queue

### 6.3.1. Operation Scenario

You can delete an apply process queue when you no longer need it.

### 6.3.2. Delete Instructions

Before deleting a Fleet, check whether the Fleet is associated with an
alias. If yes, the Fleet cannot be deleted.

### 6.3.3. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Fleet Management \> My Fleet.

3.  In the Fleet list, locate the row that contains the target Fleet,
    and click Delete.

> ![](images-en/media/image52.png)

4.  In the displayed dialog box, click OK.

# 7. Auto-Scale policy management

## 7.1. Creating an Scaling Policy

### 7.1.1. Operation Scenario

When using the application process queue, you can create and bind an
Scaling policy to control the scaling policy of the application process
queue.

### 7.1.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Fleet Management \> Scaling Policy \> Create new scaling
    policy.

> ![](images-en/media/image53.png)

3.  Set parameters such as Name, Policy Type, and Indicator. The
    following table describes the key parameters.

| **Parameters**                          | **Explained**                                                | **Example Value**        |
| --------------------------------------- | ------------------------------------------------------------ | ------------------------ |
| Name                                    | Name of the created **SCALING** policy.                      | -                        |
| Policy Type                             | Specifies the type  of the scaling policy. Currently, target-based policies are supported. | Goal-based Policy        |
| Indicator Name                          | Indicates the name  of the indicator to be controlled by the policy. Currently, the available  session percent and number are supported. | Available Session  Ratio |
| Maximum Target  Available Session Ratio | This parameter  specifies the threshold for reducing the number of application process  queues. | 60                       |
| Minimum target  available session ratio | This parameter  determines the capacity expansion threshold of the application process queue. | 30                       |

![](images-en/media/image54.png)

4\. After the parameters are set, click Create.

## 7.2. Modifying an Scaling Policy

### 7.2.1. Operation Scenario

When managing an application process queue, you can modify the Scaling
policy bound to the application process queue. Parameters that can be
modified include the name, measure name, and target value.

### 7.2.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Fleet Management \> Scaling Policy Management.

3.  In the Scaling policy list, locate the row that contains the target
    Scaling policy, and click Details.

4.  On the details page, click Modify next to the parameter, enter the
    modified content, and click Save.

![](images-en/media/image55.png)

## 7.3. Deleting an Scaling Policy

### 7.3.1. Operation Scenario

You can delete an Scaling policy that is no longer required.

### 7.3.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Fleet Management \> Scaling Policy Management.

3.  In the Scaling policy list, locate the row that contains the target
    Scaling policy, and click Delete.

> ![](images-en/media/image56.png)

4.  In the displayed dialog box, click OK.

# 8. Alias Management

## 8.1. Creating aliases

### 8.1.1. Operation Scenario

You can create aliases for the application process queue (fleet). You
can use aliasId to replace fleetId to create sessions to complete the
gray release update process. In addition, you can configure different
weights for different fleets and allocate different session creation
requests in weighted mode.

### 8.1.2. Creation Instructions

Ensure that all fleets associated with the alias are activated. The
weight of each fleet is a relative weight. The value is an integer
ranging from 0 to 100.

### 8.1.3. Operation Procedure

1.  Logging In to the GameFlexMatch Console

2.  Choose Fleet Management \> Alias Management \> Create Alias.

> ![](images-en/media/image57.png)

3. Set the required parameters.

| Parameters           | Description                                                  | Type   |
| -------------------- | ------------------------------------------------------------ | ------ |
| **aliases**          | The alias name contains 1 to 1024 characters.  It starts with a letter and consists of letters, digits, underscores (_), and  hyphens (-). | string |
| **Description**      | Description of the alias (1 to 1024  characters)             | string |
| **Type**             | Alias type (activated/deactivated). Alias of  the activated type can be used to create sessions. | enum   |
| **Message**          | Alias message, which is returned when a  session is created with an alias of the disabled type (1 - 1024 characters) | string |
| **Associated Fleet** | Select the fleet to be associated with the  alias and set the weight of the fleet. A maximum of 10 fleets can be  associated. | array  |
| **Fleet ID**         | Specifies the ID of the associated fleet. The  value is a 32-bit UUID. | string |
| **Weight**           | The value ranges from 0 to 100.                              | int    |

4.  After the parameters are set, click Create.

## 8.2. Modifying aliases

### 8.2.1. Operation Scenario

You can modify the alias name, description, type, message, and
associated fleet when the alias name, description, type, message, and
associated fleet need to be added, deleted, or modified. If the fleet
associated with the alias changes to the active state but the alias type
needs to be changed to the active state, you need to modify the alias
type.

### 8.2.2. Operation Procedure

1.  Log in to the GameFlexMatch Console

2.  Choose Fleet Management \> Alias Management and select the alias to
    be modified.

3.  Modify the value based on the site requirements. When modifying the
    associated fleet, only the weight can be modified.

## 8.3. Querying aliases

### 8.3.1. Operation Scenario

Filter aliases by alias name, type, and associated fleet.

### 8.3.2. Operation Procedure

1.  Log In to the GameFlexMatch Console

2.  Choose Fleet Management \> Alias Management.

3.  Click the search box. Search by fleet id, alias name, or alias type
    is displayed. Select a value based on the site requirements.

> ![](images-en/media/image58.png)

## 8.4. Delete Alias

### 8.4.1. Operation Scenario

If an alias is no longer used, you need to delete it.

### 8.4.2. Operation Procedure

1.  Logging In to the GameFlexMatch Console

2.  Choose Fleet Management \> Alias Management, select the alias to be
    deleted, and click Delete.

3.  In the displayed Confirm dialog box, click OK.

![](images-en/media/image59.png)

# 9. Log management

## 9.1. Creating Log Access

### 9.1.1. Operation Scenario

The GameFlexMatch platform provides the automatic management function
for instance access logs in the application process queue. You can
create a log access to record and analyze the logs of all instances.

### 9.1.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Log Management \> Create Log.

![](images-en/media/image24.png)

3.  Set log access parameters. If no log group exists, click Create next
    to Log Group Name.

![](images-en/media/image25.png)

4.  The following table lists the parameters for creating a log access.

| **Parameters**  | **To explain**                                               | **Example Value**           |
| --------------- | ------------------------------------------------------------ | --------------------------- |
| FleetId         | Creating an  application process queue to which log access belongs | -                           |
| Log Access Name | Log access name.                                             | -                           |
| Log group name. | Select the log group  to which the created log stream belongs. |                             |
| Pathway         | Log path or file.  Multiple access paths are supported,      | /root/log or  /root/log.txt |
| Description     | Description of the  log.                                     |                             |

The parameters for creating a log group are as follows:

![](images-en/media/image60.png)

| **Parameters**               | **Explained**                                                | **Example Value** |
| ---------------------------- | ------------------------------------------------------------ | ----------------- |
| Log group name.              | Log group name.                                              | -                 |
| Log retention period  (days) | Log retention period  on the LTS platform, ranging from 1 to 365 days. | 7                 |

5.  After the creation is successful, the system prompts you to create a
    log dump. If you click Yes, you will be redirected to the creation
    page. If you click No, you will skip the creation. You can create a
    log dump at any time on the log access details page.

![](images-en/media/image26.png)

6.  Note that one log access corresponds to only one log dump.

7.  After a log access task is created, it takes about 10 minutes to
    view detailed log information on LTS.

## 9.2. Creating a Log Dump

### 9.2.1. Operation Scenario

Log dumping is used to permanently dump logs temporarily stored in Log
Access to OBS buckets. After the configuration is complete, logs are
automatically dumped based on parameters.

### 9.2.2. Operation Procedure

1.  Log in to the GameFlexMatch management console.

2.  Choose Log Management \> Log Details \> Configure Log Dump.

![](images-en/media/image61.png)

3.  Enter the configuration information for creating log dump.

![](images-en/media/image62.png)
4.  Set the parameters as follows:

| **Parameters**   | **Explained**                                                | **Example Value** |
| ---------------- | ------------------------------------------------------------ | ----------------- |
| OBS bucket name  | Specifies the name  of the OBS bucket for dumping logs.      | -                 |
| Dump Path Prefix | If you enter  LTS-test, the log dump path is OBS bucket name /LogTanks/LTS-test/Y/M/D/. | -                 |
| Dump Period      | The dump period can  be 2 minutes, 5 minutes, 30 minutes, 1 hour, 3 hours, 6 hours, or 12 hours. | -                 |

5.  After the creation is successful, the OBS dump path is displayed on
    the list page. If the dump path is not configured, the OBS dump path
    is empty. The path can be linked to the OBS bucket object list page.

## 9.3. Deleting Log Access and Dump

### 9.3.1. Background

Delete the configured log access and log dump.

### 9.3.2. Operation Procedure

1.  Delete Log Access. (If log dump has been configured, delete the log
    dump first.). On the log access list page, locate the log access
    item to be deleted and click Delete.

![](images-en/media/image63.png)

2.  To delete a log dump, click Delete Log Dump in the upper right
    corner of the Log Access page.

![](images-en/media/image64.png)

# 10. Event Audit

## 10.1. Viewing the Event List

### 10.1.1. Operation Scenarios

View Game Flex Match events, including network, instance, session, and
application package events. The event levels are normal, warning, and
accident. You can configure notifications for event levels.

### 10.1.2. Operation Procedure

Choose Event Audit \> Event List to view events. You can filter events
by resource type and severity.

![](images-en/media/image65.png)

## 10.2. 11.2 Message Management

### 10.2.1. 11.2.1 Operation Scenarios

Message management manages notification subscriptions and sends
notifications to subscribers through SMN.

### 10.2.2. 11.2.2 Operation Procedure

1\. Choose Event Audit \> Message, add a topic for the event, and select
the event level for which notifications need to be sent.

![](images-en/media/image66.png)

![](images-en/media/image67.png)

2\. After a topic is created, add subscribers. The email or SMS
notification mode is supported.

![](images-en/media/image68.png)

3\. After the subscription is successfully created, SMN will send a
confirmation email. Click the link in the email to confirm the
subscription.

4\. To delete a notification, click Unsubscribe.

![](images-en/media/image69.png)

# 11. API invoking

## 11.1. Interface authentication

Authentication must be performed before all interfaces are invoked.

1\. Construct a login request. In the request, the password is encrypted
by RSA. The encrypted public key must be the same as the backend private
key deployed by FleetManager. For details about the encryption script,
see the /tools/cipher file.

POST URL: /v1/user/login

Request Body:

```json
{
    "username": "admin",
    "password": "cipher-password"
}
```

Response Body:

```json
{
    "username": "admin",
    "id": "63da950a-b3ea-11ed-bb27-fa163**********",
    "Auth-Token": "eyJhbGciO*************eXrA",
    "Activation": 1,
    "UserType": 9,
    "total_res_count": 1
}
```

2\. Add the Auth-Token field and its value in the request body header to
access the GameFlexMatch interface.

# 12. FAQs

## 12.1. Fleet and Application Package Creation

1.  How Do I Locate the Cause When Creating a Fleet Fails? During the
    creation of a fleet, the following steps are involved: synchronizing
    account information, synchronizing application package resources,
    creating a VPC, creating a subnet, creating a security group,
    creating an Scaling group, and waiting for application process
    reporting. If the creation fails, locate the fault in the following
    phase:

-   **Cause: Log in to the host where the Fleet Manager is located, open
    the /home/fleetmanager/bin/log/run/run.log file, search for the
    keyword creat_fleet, and locate the error.**

-   **Synchronize account information: Check whether the account
    information involved in GameFlexMatch is completely configured. For
    details, see the quick start.**

-   **If the application package resource synchronization fails, check
    whether the application package associated with the fleet is
    available. If the application package is unavailable, switch to an
    available one and create a new one.**

-   **Failed to create a VPC or subnet. If the VPC and subnet can be
    bound to the Fleet, check whether the bound VPC and subnet are
    available. If the bound VPC and subnet are not bound, a new VPC or
    subnet is created by default. In this case, locate the cause based
    on the returned error information in the logs.**

-   **Failed to create a security group. A security group is created by
    default during Fleet creation and can be bound to a security group.
    If the security group fails to be created, locate the fault based on
    the error information in the logs.**

-   **Failed to create the AS group. The AS component executes the AS
    group creation task. If an error occurs during this step, log in to
    the AS server and view the /home/aass/bin/log/run/run.log file to
    locate the error cause.**

-   **Waiting for application process reporting: This problem occurs
    because the application process is not reported. The main causes are
    as follows: The uniform host fails to be created (you can log in to
    the HUAWEI CLOUD console to view the creation), and the uniform host
    is created but the application fails to be started. (You need to log
    in to the uniform machine to check whether the application is
    started correctly.). The application is started normally but fails
    to be reported. (You can log in to the uniformed machine and view
    the auxproxy log in /etc/auxproxy/log/run.log to view the error
    cause.)**

2.  How can I locate the cause if the application package fails to be
    created? The application package creation process involves creating
    an ECS, pulling an application package, and creating an image. You
    need to log in to the FleetManager log to view the error cause and
    locate the fault.

-   Session creation

## 12.2. Viewing Session Error Causes

1.  Session creation failures are classified into the following types:


-   **error for timeout:**

    Check whether the 60001 port is enabled for the appgateway address.
    If the 60001 port is enabled for the appgateway address in the
    security group corresponding to the appgateway address. If the 60001
    port is not enabled, You can manually modify the ports and addresses
    opened by the security group to the AppGateway (valid for the
    fleet), and modify internal_inbound_permissions in the
    /home/fleetmanager/configmap/service_config.json file of the
    fleetmanager. The modification takes effect for the new fleet. This
    problem usually occurs during new deployment. If the problem is not
    caused by configuration problems, the network is faulty. Contact
    HUAWEI CLOUD for emergency workarounds.

     If the problem occurs occasionally, view the
     /home/appgateway/log/run/run.log log of AppGateway and the run.log log
     of AuxProxy in /etc/auxproxy/log/, trace the session life process, and
     determine whether the problem is caused by the GameFlexMatch or the
     game service.

-   Error cause: Redis or MySQL: The MySQL or Redis specifications are
    low and need to be expanded in a timely manner.

-   Error cause: The number of handles of the appgateway machine is
    incorrectly set. You need to modify the configuration.

-   there is no available process in fleet: If the number of registered
    processes in fleet is insufficient, check whether a large number of
    processes exist. If yes, manually modify the minimum number of
    instances in fleet, manually expand the capacity of the host,
    restore services in time, and then trace the cause. If this problem
    occurs occasionally, the scaling speed does not meet the traffic
    increase speed. In this case, you need to adjust the minimum number
    of instances reserved in the Scaling policy and fleet.

## 12.3. Version iteration

How to quickly and securely perform version iteration during peak hours?

GameFlexMatch can quickly update and iterate games based on alias
management. If iteration is performed during off-peak hours, you can
directly create a fleet, associate the new fleet with the alias, and
move the old fleet out of the alias. Resources can be reclaimed after
all sessions on the fleet of the old version are complete. During
version iteration during peak hours, services must not be affected. Two
update methods are available:

### 12.3.1. Step by step traffic diversion 

1.  create a fleet with the same configuration as the fleet of the old
    version. The minimum number of instances in the fleet of the new
    version must be the same as that of the old version.

2.  Associate the fleet of the new version with the alias of the old
    version. Set the weight ratio based on the proportion of the current
    number of devices.

3.  Gradually increase the weight of the fleet of the new version based
    on the distribution situation.

4.  Gradually reduce the weight of the fleet of the old version. After
    the traffic is stable, disassociate the fleet of the old version
    from the alias.

5.  After the sessions of the fleet of the old version are complete,
    reclaim the fleet resources and complete the traffic diversion.

### 12.3.2. Full traffic diversion

1.  Create a fleet with the same configuration as the fleet of the old
    version. Set the minimum number of instances of the new version
    fleet to the number of instances of the old version fleet.

2.  Enable auto scaling and associate the new version fleet with the
    alias of the old version fleet. Disassociate the fleet of the old
    version from the alias.

3.  After all sessions on the fleet of the old version are complete,
    change the minimum number of instances of the fleet of the new
    version to the minimum number of instances set by the fleet of the
    old version. Then, reclaim the resources of the fleet of the old
    version.

We recommend using the second change method because that can be
completed quickly and securely during peak hours.

### 12.3.3. Other questions

1.  Is the CPU or memory abnormal on the uniform machine?

> Log in to the uniformed machine and check whether the number of game
> processes is abnormal. If yes, check whether a single game process
> occupies a large number of CPU or memory resources and adjust the
> number of processes started on a single machine.

2. Is the resource monitoring exception of GameFlexMatch?

> Check the current service traffic information. If the traffic
> increases sharply compared with the previous period, expand the
> capacity or increase the specifications of the nodes in a timely
> manner. If the traffic is normal but the CPU or memory alarm persists,
> check whether a large number of abnormal sessions are displayed on the
> console home page. If yes, rectify the fault based on the session
> creation failure type.
