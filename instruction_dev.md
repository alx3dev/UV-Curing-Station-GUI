# Instructions for reading the source code

 - **app** - application logic directory
 - **theme** - resources directory
 - **translation** - directory with translated words

Application logic is divided into multiple files in a single directory:   

 - Files that start with `uvs_` handle app logic.  
 - Files that start with `screen` handle window logic.  
 - `helper_bindings` handle fyne bindings, to make screen code more readable.  
 - `translate.go` handle translation logic.  
 - `translate/strings.go` handle languages.  

Notifications allow sending system notification immediately, or prepare for another start.  
If multiple notifications are prepared, they will be sent as a single notification.