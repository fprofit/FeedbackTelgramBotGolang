# Feedback Telgram Bot Golang
### Create a file "settings.txt" in the folder of the project or binary file with contents

    { 
	    "botToken":"YOUR_BOTTOKEN",
	    "adm_id":YOUR_ID (int),
	    "text":{"Lang_Code":"Response text in language == "lang_code" to the /start command",
	    		"default":"Command response text /start"
	    	   }
    }
    
* botToken - string
* adm_id - int
* text - map[string]string (the "default" key is required)


### Select a message to reply to it
### To get information about the sender, reply with a command to the selected message "/user_info" (Data may not be available due to profile privacy)