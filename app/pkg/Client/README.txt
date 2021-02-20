Client Service

The client service is in charge of communicating with the clients as well as managing the data stored about them. It will communicate with
the Persistence and Broker repositories and can:
    - Add new clients
    - Delete clients
    - Manage clients

It will be accessible over two interfaces:
    - Interface 1: (Is used to initiate communication and store information regarding a client)
        * AddClient
        * ConfigureClientBackupList
        * ClientDirectoryScan
        * DeleteClient
        * AddClientPolicy
        * DeleteClientPolicy
    - Interface 2: (Is used for communication FROM the client device and often responds to presentation layer)
        * DirectoryScanResult(Directory[])
        * ConfigRequest(clientName)
        * NewClient(config)

Client Service is in charge of the following:
    - Adding new client devices to the system
    - Updating client devices. This could be for the following reasons:
        * When they have been added to a new policy. (add bind to queue)
        * When they have been removed from a policy. (remove bind from queue)
        * When the user changes what needs to be backed up on the client and calls for a directory scan
    - Managing and storing information from client devices in a repository


AddClient
- To add a client the AddClient method is called. The following process will then happen:
    * Broker method NewClientDetails is triggered by AppClient
        - This tells repository to communicate with the client and request details regarding the client
    * ClientService recieves a response through the NewClient method on interface 2
        - AddClient will then respond to the caller passing through the client found
        - The Client is then added to the database.

ConfigureClientBackupList(client, directoryTree)
- This receives a directory tree. 
- It then takes this tree and calls UpdateClientBackupList on the presistent repository


ClientDirectoryScan(clientName, path)
- This is in charge of the directories on a client. The process is as follows:
    * The client name is passed through, as well as a path. 
    * A client DirectoryScan on broker repository is requested passing through the path
        - The directory scan only scan's 4 layers at a time, this is to speed things up
    * ClientService then receives the result through DirectoryScanResult which in turn passes it back

DeleteClient(clientName)
- Takes a client name and performs DeleteClient on persistent repository

AddClientPolicy(clientName, policy)
- This is triggered on the back of EditPolicy in the Policies service.
- It takes the client name and makes call on persistent repository to add the policy(passed through) to the client's record
- Will then trigger AddPolicy on the broker repository passing through the policy and client name

DeleteClientPolicy(clientName, policy)
- This is triggered on the back of DeletePolicy in the Policies service.
- It takes the client name and makes call on persistent repository to delete the policy(passed through) from the client's record
- Will then trigger DeletePolicy on the broker repository passing through the policy and client name
