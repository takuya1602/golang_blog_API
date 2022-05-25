## API

| Objective                            | Endpoint                  | Method |
| ------------------------------------ | ------------------------- | ------ |
| List of parent-categories            | /parent-categories        | GET    |
| Add parent-category                  | /parent-categories        | POST   |
| Change the parent-category name      | /parent-categories/:slug  | PUT    |
| Delete the parent-category           | /parent-categories/:slug  | DELETE |
| List of categories                   | /categories               | GET    |
| Add category                         | /categories               | POST   |
| Change the category name             | /categories/:slug         | PUT    |
| Delete the category                  | /categories/:slug         | DELETE |
| List of posts                        | /posts                    | GET    |
| List of posts of the parent-category | /posts?parent-category={} | GET    |
| List of posts of the category        | /posts?category={}        | GET    |
| Add post                             | /posts                    | POST   |
| Update the post                      | /posts/:slug              | PUT    |
| Delete the post                      | /posts/:slug              | DELETE |

tet