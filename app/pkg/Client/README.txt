Client Service

The client service is in charge of communicating with the clients as well as managing the data stored about them. It will communicate with
the Persistence and Broker repositories and can:
    - Add new clients
    - Delete clients
    - Manage clients

It will be accessible over two interfaces:
    - Interface 1: (Is used to initiate communication regarding a client. Used by GUI etc.)
        * AddClient
        * ConfigureClientBackupFiles
        * DeleteClient
        * UpdateClientPolicy
        * RemoveClientPolicy
    - Interface 2: (Is used for communication FROM the clients)
        * FilescanResult(filescan)
        * ConfigRequest(clientName)
        * NewClient(config)

clientService is in charge of the following:
    - Adding new clients to the system
    - Updating clients. This could be for the following reasons:
        * When they have been added to a new policy. (add bind to queue)
        * When they have been removed from a policy. (remove bind from queue)
        * When the user changes what needs to be backed up on the client and calls for a filescan


Adding a client
- To add a client the AddClient method is called. The following process will then happen:
    * Broker method NewClientDetails is triggered by AppClient
        - This tells repository to communicate with the client and request details regarding the client
    * Client recieves a response through the NewClient method on interface 2
        - Broker method NewClientFilescan is triggered
    * Client receives a response through FilescanResult on interface 2
        - The Client is then added to the database. 
        - This response is returned to the caller of AddClient

ConfigureClientBackupFiles(clientName)
- This manages the files the client will backup of and can occur when a client is added or when modifying the client
The process is as follows:
    * the client is retrieved from the persistence repositories
    * 

    