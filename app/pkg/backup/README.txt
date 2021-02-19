Backup Service

backupService will be in charge of performing and creating backups. It will communicate with the Persistence and Broker repositories and can:
    - Start Backup Run
    - View Backup Runs

There are two interfaces:
    - Interface 1
        * TriggerBackupRun(Policy Name)
        * ViewBackupRuns()
    - Interface 2
        * BackupExternalCommunication (Returns: a Channel)
            * The channel is used to open up a connection between the ExternalConsumer



TriggerBackupRun(Policy Name)
    - Takes a Policy Name as input
    - Returns a Channel
This is the largest and most complex function as it will run and manage the backup process. It is called by either the user manually through the UI or the queryBackupService calls the TriggerBackupRun method
The steps are as follows:
    * Query PersistenceRepository for information regarding the inputted policy
        - Type of backup to be performed
        - The clients to run
    * After this it will then call the BrokerRepository BackupRun Interface function, passing through the policy name.
    * When the BackupExternalCommunication method is called for the first time it returns a channel. This opens up communication between the RabbitConsumer and the service.
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

        