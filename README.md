# CLI

| Command | Code | 
| ---------|:--------:|
| build  | ```skaffold build```  |
| render  | ```skaffold render```  |
| run  | ```skaffold run```  |
| test  | ```skaffold test```  |
| verify  | ```skaffold build -q \|skaffold verify --build-artifacts - ```  |
| debug | ```skaffold build -q\|skaffold deploy -v debug --build-artifacts - 2>&1 \| grep -i 'Running command' ```|
| delete | ```skaffold delete``` |

## Run with *Dev* profile

| Command | Code | 
| ---------|:--------:|
| build  | ```skaffold build -p dev```  |
| render  | ```skaffold render -p dev```  |
| run  | ```skaffold run -p dev```  |
| test  | ```skaffold test -p dev```  |
| verify  | ```skaffold build -p dev -q \|skaffold verify -p dev --build-artifacts - ```  |
| delete | ```skaffold delete -p dev``` |

