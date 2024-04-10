Seller
id, uuid, name, email, created_at, updated_at, deleted_at

User
id, uuid, name, mobile_number, email, address_id
created_at, updated_at, deleted_at

addresses
id, user_id, address, pincode, city, state, is_default

Deal
id, uuid, StartTime, EndTime, NoOfMaxItems, Price, NoOfItemsAvailable

Order
user_id, deal_id, status, created_at, updated_at, deleted_at

Api EndPoints

Seller
    /login
    /createDeal
    /updateDeal
    /endDeal


User
    /login
    /deals - view all deals
    /order - claims a deal
    



TODO:
Add Payment Table to track order placement
User, Seller Login
