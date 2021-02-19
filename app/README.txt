 - Use an Event Bus for communication between services.
 - Use a Command/Query Bus to communicate to services 


 DevicesHandler
  - Is a RabbitMQ consumer that is called at the start of the service. 
  This has a controller that implements the following Interfaces: 
        - ClientIntrerface
        - BackupInterface
        - RestorInterface


