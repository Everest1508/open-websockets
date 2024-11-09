# Open Websockets 

This project is a Go-based WebSocket application using the Gorilla WebSocket library. It supports real-time, bidirectional communication using a Pub/Sub (Publish/Subscribe) pattern.

## Connecting with Query Parameters

### Room and Role Parameters

The WebSocket server accepts two query parameters upon connection:

- **`room`**: Specifies the room or channel the client is joining. Rooms allow for message separation, where clients in different rooms won’t receive each other's messages.
  
- **`role`**: Specifies the role of the client, which can be either:
  - **`publisher`**: Clients with this role can publish messages to the room they are connected to.
  - **`subscriber`**: Clients with this role can only receive messages from the room they are connected to.

### Example Connection URLs

- **Publisher**: Connects to room `example-room` with the ability to publish messages.
    ```
    ws://localhost:8080/ws?room=example-room&role=publisher
    ```

- **Subscriber**: Connects to room `example-room` to only receive messages.
    ```
    ws://localhost:8080/ws?room=example-room&role=subscriber
    ```

Using these parameters, clients can join specific rooms and take on either publishing or subscribing roles, enabling a flexible Pub/Sub WebSocket experience.
