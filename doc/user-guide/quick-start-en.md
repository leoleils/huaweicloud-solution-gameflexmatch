# GameFlexMatch Quickly Start: Content 
- [GameFlexMatch Quickly Start: Content](#gameflexmatch-quickly-start-content)
- [1. Managing users and associating resource users](#1-managing-users-and-associating-resource-users)
  - [1.1. Complete information](#11-complete-information)
- [2. Application upload and image creation](#2-application-upload-and-image-creation)
  - [2.1. Creating an Application Package](#21-creating-an-application-package)
  - [2.2. View the application package details and check the application status.](#22-view-the-application-package-details-and-check-the-application-status)
- [3. Creating an Instance Flavor](#3-creating-an-instance-flavor)
  - [3.1. Operation Procedure](#31-operation-procedure)
- [4. Process for creating a fleet](#4-process-for-creating-a-fleet)
  - [4.1. Creating a Fleet](#41-creating-a-fleet)
  - [4.2. Check whether the Fleet is successfully created.](#42-check-whether-the-fleet-is-successfully-created)
  - [4.3. (Optional) Create an Scaling policy and enable the Scaling function.](#43-optional-create-an-scaling-policy-and-enable-the-scaling-function)
- [5. (Optional)Creating an alias association](#5-optionalcreating-an-alias-association)
  - [5.1. Create Apply Process Queue Alias](#51-create-apply-process-queue-alias)
- [6. (Optional) Creating Log Ingestion and Dump](#6-optional-creating-log-ingestion-and-dump)
  - [6.1. Creating Log Access](#61-creating-log-access)

# 1. Managing users and associating resource users

- There are two types of users of GameFlexMatch,
    + **Console login user**: Used to log in to the GFM console for routine management and O&M.
    + **HUAWEI CLOUD tenant**: used to create and manage computing and network resources.
  Both of them need to be configured.

- When logging in to the GameFlexMatch console for the first time, the
    administrator or common user needs to reset the password and
    associate the password with the HUAWEI CLOUD tenant so that the
    GameFlexMatch console can work properly.

## 1.1. Complete information

1.  Log in to the console using the administrator account and password.
    If the administrator logs in to the console for the first time,
    change the password and log in again.

![](images-en/media/image2.png)

2.  If no resource tenant is associated at the first login, the
    administrator needs to associate the resource tenant first.

![](images-en/media/image3.png)
![](images/media/add-user.png)  

3.  **Enter information about associated resource tenants.**

Use tenant information:

![](images-en/media/image4.png)

-   **Tenant name**: The name of the tenant bound to the GameFlexMatch platform is also your HUAWEI CLOUD tenant name. You can log in to [IAM-users](https://account-intl.huaweicloud.com/usercenter/#/iam/users) and select a HUAWEI CLOUD tenant.

-   **Project ID**: Project ID of the API used by GameFlexMatch to create resources. To obtain the project ID,  You can log in to [mine-apiCredential](https://console-intl.huaweicloud.com/iam/#/mine/apiCredential) to select the ID of the regional project that you want to bind.

-   **Access Key**: Key used by GameFlexMatch to access APIs. You can obtain it from HUAWEI CLOUD console -> My Credentials -> [Access key](https://console-intl.huaweicloud.com/iam/#/mine/accessKey).

-   **Secret Access Key**: Key used by GameFlexMatch to access APIs. You can obtain it from HUAWEI CLOUD console -> My Credentials -> [Access key](https://console-intl.huaweicloud.com/iam/#/mine/accessKey).

-   **Region**: Resource tenants are used to create regions for GameFlexMatch resources. The information must correspond to the project ID.

-   **Key name**: Login authentication key pair for GameFlexMatch elastic VM capacity expansion. You can use this key pair to log in to the VM of the battle server. You can log in to the DEW console and choose [Key Pair Management](https://console-intl.huaweicloud.com/console/#/dew/kps/kpsList/accountKey) to view the account key pair. If no key pair is available, create one.

-   **Cloud Service Agency Name**: You can choose Identity and Access Management > [Agency] (https://console-intl.huaweicloud.com/iam/#/iam/agencies) on the HUAWEI CLOUD console to view the ICAgent and LTS log dump.
    1.  You can customize the agency name, for example, gfm-lts-agency.
    2.  Select **Cloud Service** for Agency Type.
    3.  Cloud Service: Select ECS.
    4.  Click Next and grant the agency permissions:'APM FullAccess' and'LTS FullAccess'.
    5.  After the agency is created, you can configure the agency name'gfm-lts-agency' to the corresponding location.

Application package information (optional. If this parameter is left blank, the default value is used.)

-   **Auxproxy Path**: OBS path for storing the Auxproxy file. Ensure that **auxproxy.zip** has been uploaded to the OBS path to be configured.
-   **ECS-Application Package Configuration Script Path**: OBS path for storing the script file used to create the ECS resource application package. Ensure that **image_env.sh** has been uploaded to the OBS path to be configured.

-   **Region**: Region where the OBS bucket resides

-   **Container-Application Package Configuration Script Path**: OBS path for storing the script for creating a containerized application package; If a container is used, this parameter is mandatory. Ensure that **docker_image_env.sh** has been uploaded to the OBS path to be configured.

The resource tenant is successfully associated and can use
GameFlexMatch.

# 2. Application upload and image creation

-   The application image contains the AuxProxy service component and
    interconnected backend service applications, and related startup
    configuration has been completed. The template used for elastic VM
    expansion has been configured.

-   Prerequisites:

> Ensure that the resource tenant has configured correct tenant
> information for the management tenant.
>
> Ensure that the uploaded server applications can be started normally.

For details, see the application package part in the guide.

## 2.1. Creating an Application Package

1.  User login console

2.  Go to the Application Package Management page and click "Create App
    Package" to create an application.

> ![](images-en/media/image5.png)

3.  Enter the image creation information and click Create
    Server-Application.

![](images-en/media/image6.png)

## 2.2. View the application package details and check the application status.

1.  On the application package management page, find fake-server.

> ![](images-en/media/image7.png)

2.  Click to view details and check whether the application status is
    READY.

> ![](images-en/media/image8.png)

3.  When the status is Ready, you can use the application to create a
    fleet.

# 3. Creating an Instance Flavor

-   An ECS flavor group provides a group of ECS flavors. You can select
    the number of CPU cores and memory size based on service
    requirements.

-   Prerequisites:

> The user tenant information must be correctly configured.

## 3.1. Operation Procedure

1.  A user logs in to the console.

2.  Choose Configuration \> Instance Flavor Group and create a new
    flavor group.

> ![](images-en/media/image9.png)

3.  Create an instance specification. For example, the 2U4G instance
    specification of the VM type is used. Enter the instance
    specification group name, select the instance specification, and
    click OK.

> ![](images-en/media/image10.png)

# 4. Process for creating a fleet

-   An application process queue (fleet) is a queue that manages backend
    service application clusters. The number of backend service
    applications can be manually or automatically increased to meet
    different load requirements,

-   Prerequisites

> An application package in the Ready state has been created.
>
> You need to create an instance specification group.

## 4.1. Creating a Fleet

1.  User login console

2.  The page for creating a fleet is displayed.

> ![](images-en/media/image11.png)

3.  Enter the fleet creation information. JSON files can be imported.
    For details about the parameters, see the application process queue
    management section.

4.  Click Create to create a fleet.

> ![](images-en/media/image12.png)

## 4.2. Check whether the Fleet is successfully created.

1.  Go to the fleet list and find the newly created fleet.

> ![](images-en/media/image13.png)

2.  Click Details to view the fleet information.

> ![](images-en/media/image14.png)

3.  If the fleet status changes from Creating to Active, the creation is
    successful. If the fleet status changes from Creating to Abnormal,
    the creation fails. You can locate the fault in the details page.
    For details, view the FleetManager service logs. For common errors,
    see Error Information.

## 4.3. (Optional) Create an Scaling policy and enable the Scaling function.

-   Before performing this step, ensure that the fleet for which the
    Scaling policy is created is activated.

-   Currently, only the available session ratio can be selected as the
    benchmark. Computing resources can be scaled in or out based on the
    maximum number of sessions that can be carried, the current session,
    and the dynamic relationship between the available session ratio.

-   Prerequisites:

> The Fleet to be bound is in the Active state.

1.  Find the entry for creating an Scaling policy.

> ![](images-en/media/image15.png)

2.  Select an Scaling policy, set parameters, and click Create.

> Note: Rate of available sessions = (Maximum number of sessions -
> Number of used sessions)/Maximum number of sessions
>
> ![](images-en/media/image16.png)

3.  On the Scaling policy details page, click Create Association.

> ![](images-en/media/image17.png)

4.  Select the fleet to be associated with the policy.

> ![](images-en/media/image18.png)

5.  Check whether the Scaling policy is successfully created.

> ![](images-en/media/image19.png)

6.  The details about the fleet bound to the Scaling policy are
    displayed.

> ![](images-en/media/image20.png)

7.  Modify Basic Information \> Enable Auto Scaling to enable the auto
    scaling capability.

> ![](images-en/media/image21.png)

8.  Now, GameFlexMatch can flexibly scale computing resources based on
    the load.

# 5. (Optional)Creating an alias association

-   The fleet alias supports dark launch. Multiple fleets can be
    associated with the same alias. When creating a session, alias_id
    can be carried instead of fleet_id to create a session. Different
    fleets have different weights and different session creation
    requests can be weighted.

-   Prerequisites:

> Ensure that the associated fleet is in the Active state.

## 5.1. Create Apply Process Queue Alias

1.  The page for creating an alias is displayed.

> ![](images-en/media/image22.png)

2.  Enter the alias creation information.

> ![](images-en/media/image23.png)

3.  Click Create to create an alias. Now you can use the alias to create
    a session.

# 6. (Optional) Creating Log Ingestion and Dump

-   Log Access allows you to temporarily store application logs
    generated by instances in Scaling group to LTS, analyze and record
    the logs, or permanently dump the logs to OBS.

-   Prerequisites:

> You need to configure the cloud service agency in the tenant
> information, correctly authorize the cloud service agency, and pack
> the application package (ensure that the ICAgent is correctly
> installed in the image).
>
> The Fleet to be configured must be in the Active state.

## 6.1. Creating Log Access

1.  On the Log Management page, click New Log.

> ![](images-en/media/image24.png)

2.  Enter information about log access. You need to select a log group
    for log access. If there is no log group, create a log group.

> ![](images-en/media/image25.png)

3.  After the creation is successful, you can configure the automatic
    transfer of logs to OBS. You can also select No to skip this step
    and create a log at any time on the log details page.

> ![](images-en/media/image26.png)

4.  If you select Create log transfer, set the parameters for creating
    log transfer and click Create.

![](images-en/media/image27.png)

5.  After the creation is successful, a new record is displayed on the
    Log Management page. You can click the log stream or OBS path to go
    to the corresponding HUAWEI CLOUD service console.
    ![](images-en/media/image28.png)