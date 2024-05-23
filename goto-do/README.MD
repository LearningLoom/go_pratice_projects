# Todos app 

## this is a special todo app, this app saves the todos in Json file called todos.json 
- with this app you can 
    - create a todo 
    - List all todos 

## this app used ``` os.Writefile() and os.ReadFile()`` functions 
## also we have created a custom package called todos and haved used in our main.go 

## other functionality 
- if the todos.json does not exists this app will create it
- it wil check if the todos file is empty with the help of this function 
    ```go 
    func isJSONFileEmpty(filePath string) (bool, error) {
        info, err := os.Stat(filePath)
        if err != nil {
            return false, err
        }

        // Check file size
        if info.Size() == 0 {
            return true, nil
        }

        // Read file content
        data, err := os.ReadFile(filePath)
        if err != nil {
            return false, err
        }

        // Check if it's an empty JSON object or array
        var obj interface{}
        err = json.Unmarshal(data, &obj)
        if err != nil {
            return false, err
        }

        switch obj.(type) {
        case map[string]interface{}:
            return len(obj.(map[string]interface{})) == 0, nil
        case []interface{}:
            return len(obj.([]interface{})) == 0, nil
        default:
            return false, fmt.Errorf("unexpected JSON type")
        }
    }
    ``` 
- and if it is empty it will create an empty array json file with the help of this function 
    ```go
    func createJsonEmptyArrayFile(filePath string) {
        var emptyArray []interface{}

        jsonData, err := json.Marshal(emptyArray)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        err = os.WriteFile(filePath, jsonData, 0775)
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }
    ```
- all these functions are written in helperFunctions.go file in todos package 

### A Linux binary is also included in the codebase with the name of go_to_do

# Thank you.
