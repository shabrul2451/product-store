## Product Store

#### Clone Repository
```bash
 git clone https://github.com/shabrul2451/product-store.git
```
#### Change into the project directory

```bash
cd product-store
```

#### Run the application
```bash
go run main.go
```


#### Run the LRU problem
```bash
go run problem_solving/main.go
```

#### Open your web browser and visit http://localhost:8085 to access the music streaming service.

## Api Endpoints
1. #### CRUD Operations for Brands
- List Brands (GET): `/api/v1/brands`
- Get Brand by ID (GET): `/api/v1/brands/:id`
- Create Brand (POST): `/api/v1/brands`
```json
{
  "name": ""
}
```
- Update Brand (PUT): `/api/v1/brands/:id`
```json
{
        "name": "",
        "status_id": 1
}
```
- Delete Brand (DELETE): `/api/v1/brands/:id`

2. #### Api Endpoints for Products
- List Products (GET): `/api/v1//products/by_filters?defined_params=`
```
## Params
name=
category=
supplier=
max_price=
min_price=
verified_supplier= {values should be true/false)
brands={use "," for multiple brands. eaxample:- brand-1, brand-2}
```
- Create Product (POST): `/api/v1/products`
```json
{
            "name": "",
            "description": "",
            "specifications": "",
            "brand_id": "",
            "category_id": "",
            "supplier_id": "",
            "unit_price": 0,
            "discount_price": 0,
            "tags": ""
}
```

3. #### Api Endpoints for Categories Tree
- List Products (GET): `/api/v1//categories/tree`

### Response Format
- #### Success response
```json
{
  "status": 0,
  "message": "operation successful",
  "data": []
}
```

- #### Error response
```json
{
  "status": 0,
  "message": "operation successful"
}
```

- #### Pagination response
```json
{
  "status": 0,
  "message": "",
  "data": [],
  "total": 0,
  "page": 0,
  "pageSize": 0,
  "hasMore": false
}
```