Backup Service

backupService will be in charge of performing and creating backups. It will communicate with the Persistence and Broker repositories and can:
    - Create policies
    - Edit policies
    - Delete polices
    - Run Policies

The Interfaces are as follows:
    - CreatePolicy
    - EditPolicy
    - DeletePolicy
    - TriggerPolicyRun
    - BackupChannel (Returns: a Channel)
        * The channel is used to open up a connection between the RabbitConsumer


Create Policies
 - When a policy is created it will take the following attributes:
    * Name (String)
    * Clients (List)
    * Retention_Time (Integer)
    * Runs (List)
    * Type - Two options (String): 
        - Full: Just full backups. No Incremental
        - Incremental: Both Incremental and Full
    * Scale - Scale can be set to:
        - Day: If set to Day then the schedules below will be determined in Hours
        - Week: If set to Week then the schedules bellow will be determined in Days
        - Month: If se to Month then the schedules bellow will be determined in Weeks
    * Full_Backup_Schedule - Lists out the Full schedule as determined by the scale
    * Integer_Backup_schedule - Lists out the Incremental the schedule as determined by the scale
- After the creation of the policy it will be inputted into the Database via Persistence repository
- Once inputted in database an event will be called to the ClientManagment service wich will inform the relevent clients to update route_key_bindings to new policy

Edit Policies
 - Similar to creation but will contain the following:
    * Policy Name
    * List of attributes to be changed and the changes in them. 
- If clients have been added/removed then the ClientManagment service will be informed

Delete Policies
- Will take the Policy Name and do the following:
    * will query and retrieve from the database the clients associated with the policy.
    * Will create an event for the ClientManagment service to remove the route_key_bindings for that policy
    * Will delete the policy from the database

Run Policies
This is the largest and most complex function as it will run and manage the backup process. It is called if either the user manually or the queryBackupService calls the TriggerPolicyRun method
The steps are as follows:
    * Retrieve the following from thePersistenceRepository:
        - Type of backup to be performed
        - The clients to run
    * After this it will then call the BrokerRepository BackupRun Interface function, passing through the policy name.
    * When the BackupChannel method is called for the first time it returns a channel. This opens up communication between the RabbitConsumer and the service.
    * Models are initialised to hold data:
        - ClientFile struct: There is one for each file. These are created upon receipt of Client Job Message. Initialy placed in ClientJob map. 
        - ClientRun sturct: There will be one per client in the policy. Will hold ClientFile structs that have been confirmed, as well as a status, which is initially set to false
        - BackupRun struct: Will hold ClientRun structs.

    * When the service recieves data from the channel it is of the following types:
        - Client Job Message: Holds all the files that the client will backup.
        - Client File Success Message: Confirming that the client has copied a file and is sending it two StorageNode
        - Client File Failure Message: Message confirming that the client couldn't backup the file. 
        - Storage Node File Success Message: Confirming that the StorageNode has received and saved the data.
        - Storage Node File Failure Message: Confirming that the StorageNode couldn't save the file.

        Client Job Message
            - This is stored in the ClientFile struct. This is made up of the following:
                * Unique ID of the file
                * Status
                * Filepath
                * SNFilePath
                * Filename
                * Permissions
                * Checksum
            - Upon receipt of the message the following is supplied in the message: Unique ID, Filepath, Filename, Permissions, Checksum. These are then set
            
        
        Client File Success Message:
            - The Unique ID is supplied as reference. This is then matched against the ClientFile struct. The status is then set to "Client Confirmation"

        Client File Failure Message:
            - The Unique ID is supplied as reference. This is then matched against the ClientFile struct. The status is then set to "failed". It is then removed frim the ClientJob map and added to the ClientRun struct

        Storage Node File Success Message:
            - The Unique ID is supplied as reference. This is then used to find the file in the ClientFile struct. The SNFilepath is then added and the status changed to "Successful".
            - Whilst the ClientJob map has more than one ClientFile the successful ClientFile struct is moved from the ClientJob map and placed in the ClientRun struct. 
            - When there is only one ClientFile in the ClientJob it is marked as succesfull and placed in the ClientRun struct.
            - If more than one the statuses in each ClientRun struct are marked as false then the policy is still in operation, in which case the clientRun struct is then marked as true.
            - However, if the the ClientRun struct that the ClientFile is part of is the only one marked as false, then then it is marked as true and the policy run is complete. The channel will then be closed.

         Storage Node File Failure Message:
            - The Unique ID is supplied as reference. This is then matched against the ClientFile struct. The status is then set to "failed". It is then removed frim the ClientJob map and added to the ClientRun struct

        