---
markmap:
  colorFreezeLevel: 2
  maxWidth: 300
---

# Stickers Go API

## Users
- Endpoints about the user management. `/sign-in` returns the JWT token if credentials (email and password) match.

### **[POST]** `/sign-up`

- Hash Password
- Save into users table

### **[POST]** `/sign-in`

- Get user by email
  - ==email is unique==
- Check user password
- Generate jwt token

## Stickers

- Endpoints to manage stickers. Each endpoint is protected by an Authorization Token (JWT).

### **[POST]** `/stickers`
- Check JWT token
- Save it into stickers table

### **[GET]** `/stickers`
- Check JWT token
- Select stickers of user_id
- Return the list of stickers

### **[GET]** `/stickers/:id`
- Check JWT token
- Select the sticker of user_id
- Retrieve the sticker details by its id

### **[PUT]** `/stickers/:id`
- Check JWT token
- Verify if the sticker to be updated belongs to the user of the request
- Update it into stickers table

### **[DELETE]** `/stickers/:id`
- Check JWT token
- Verify if the sticker to be updated belongs to the user of the request
- Delete it from stickers table
