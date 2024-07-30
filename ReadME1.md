#Then build the image:

- docker build -t server_new:latest -f Dockerfile.server_new .
- docker build -t client_new:latest -f Dockerfile.client_new .





#Push to dockerhub or container registry

- docker push debolek/client_new:latest
- docker push debolek/server_new:latest

  
<img width="1280" alt="Screenshot 2024-07-30 at 03 52 22" src="https://github.com/user-attachments/assets/73db2850-86a1-487d-b587-875b4a5649e7">




![immagine](https://github.com/user-attachments/assets/6d70e163-ee89-46a3-9149-fb2fb011f48d)



<img width="1280" alt="Screenshot 2024-07-30 at 03 53 32" src="https://github.com/user-attachments/assets/21f7ab55-1eff-48fc-82d4-fdc551436301">

  

