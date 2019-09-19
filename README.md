# IOSIF

IOSIF is simple message queue. Created for architectures where asynchronous requests are used.  


####Tutorial

* Create topic
    #####Request  
    * Method `GET`  
    * Query  `topic=<topicName>`  
    * URL   `/topic`
* Subscribe
    #####Request
   * Method `POST`  
   * Queries 
        * `topic=<topicName>`
        * `autoCounter=true` if cursor should move automatically    
   * URL   `/subscribe`
* UnSubscribe
    #####Request
   * Method `GET`  
   * Headers 
        * `x-sub-token=<token>`    
   * URL   `/unsubscribe`
    #####Response
    * Type JSON  
    `{"topic": "75e22153-e207-4be6-a58c-f5a09e4de36c"}`
* Push message
    #####Request
   * Method `POST`  
   * Query  `topic=<topicName>`      
   * URL   `/`
* Get message
    #####Request
   * Method `GET` 
   * Query  `topic=<topicName>`       
   * Headers 
        * `x-sub-token=<token>`    
   * URL   `/`
    #####Response
    * Type JSON    
    ```{
         "trace_id": "b26d8035-fbef-4d9d-a4ad-4e1a4e0f699d",  
         "Index": 0,  
         "topic": "test",  
         "time_stamp": "2019-09-19T16:42:50+03:00",  
         "key": "test key",   
         "value": "test value"  
    }      