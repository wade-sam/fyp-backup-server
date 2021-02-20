PolicyService
Policy service is in charge of manging and showing everything todo with policies.
It will communicate with the persistence and broker repositories and can:
    - Create policies
    - Edit policies
    - Delete polices

The Interface properties are as follows:
    - CreatePolicy
    - EditPolicy
    - DeletePolicy

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
        - Month: If set to Month then the schedules bellow will be determined in Weeks
    * Full_Backup_Schedule - Lists out the Full schedule as determined by the scale
    * Incremental_Backup_schedule - Lists out the Incremental schedule as determined by the scale
- After the creation of the policy it will be inputted into the Database via Persistence repository
- Once inputted in database an event will be called to the ClientService wich will inform the relevent clients to update route_key_bindings for the new policy

Edit Policies
 - Similar to creation but will contain the following:
    * Policy Name
    * List of attributes to be changed and the changes in them. 
- If clients have been added/removed then a relevent event will be created for the ClientService

Delete Policies
- Will take the Policy Name and do the following:
    * will query and retrieve from the database the clients associated with the policy.
    * Will create an event for the ClientService to remove the route_key_bindings for that policy
    * Will delete the policy from the database
