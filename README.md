# IOSIF
{::nomarkdown}  
svg id="Component_1_1" data-name="Component 1 – 1" xmlns="http://www.w3.org/2000/svg" width="90" height="62" viewBox="0 0 90 62">
  <g id="Rectangle_1" data-name="Rectangle 1" fill="#fff" stroke="#707070" stroke-width="1">
    <rect width="90" height="62" stroke="none"/>
    <rect x="0.5" y="0.5" width="89" height="61" fill="none"/>
  </g>
  <line id="Line_1" data-name="Line 1" x2="18" transform="translate(17.5 22.5)" fill="none" stroke="#707070" stroke-width="1"/>
  <g id="Ellipse_1" data-name="Ellipse 1" transform="translate(59 13)" fill="#fff" stroke="#707070" stroke-width="1">
    <circle cx="9" cy="9" r="9" stroke="none"/>
    <circle cx="9" cy="9" r="8.5" fill="none"/>
  </g>
  <path id="Union_3" data-name="Union 3" d="M329,63.5h0ZM332.5,50,336,63.5Z" transform="translate(-285.743 -31.873)" fill="rgba(0,0,0,0)" stroke="#707070" stroke-linejoin="round" stroke-width="1"/>
  <path id="Union_2" data-name="Union 2" d="M393,50l-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm-4,0-4,6Zm0,0h0Z" transform="translate(-316.5 -11.5)" fill="none" stroke="#707070" stroke-linejoin="round" stroke-width="1"/>
</svg>  
{:/}


IOSIF is simple message queue. Created for architectures where asynchronous requests are used.  

![](https://img.shields.io/github/issues/KoDDrovosek/IOSIF) 
![](https://img.shields.io/github/release-date-pre/KoDDrovosek/IOSIF) 
![](https://img.shields.io/github/stars/KoDDrovosek/IOSIF) 

#### Endpoints

* Create topic
    ##### Request  
    * Method `GET`  
    * Query  `topic=<topicName>`  
    * URL   `/topic`
* Subscribe
    ##### Request
   * Method `POST`  
   * Queries 
        * `topic=<topicName>`
        * `autoCounter=true` if cursor should move automatically    
   * URL   `/subscribe`
* UnSubscribe
    ##### Request
   * Method `GET`  
   * Headers 
        * `x-sub-token=<token>`    
   * URL   `/unsubscribe`
    ##### Response
    * Type JSON  
    ```
    {
        "topic": "75e22153-e207-4be6-a58c-f5a09e4de36c"
    }
* Push message
    ##### Request
   * Method `POST`  
   * Query  `topic=<topicName>`      
   * URL   `/`
* Get message
    ##### Request
   * Method `GET` 
   * Query  `topic=<topicName>`       
   * Headers 
        * `x-sub-token=<token>`    
   * URL   `/`
    ##### Response
    * Type JSON    
    ```
    {
         "trace_id": "b26d8035-fbef-4d9d-a4ad-4e1a4e0f699d",  
         "Index": 0,  
         "topic": "test",  
         "time_stamp": "2019-09-19T16:42:50+03:00",  
         "key": "test key",   
         "value": "test value"  
    }      