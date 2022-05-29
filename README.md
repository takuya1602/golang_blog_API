# go-blog-api



|                                           | Endpoint                                      | Method |
| ----------------------------------------- | --------------------------------------------- | ------ |
| get all categories                        | /categories                                   | GET    |
| add category                              | /categories/:slug                             | POST   |
| update category                           | /categories/:slug                             | PUT    |
| delete category                           | /categories/:slug                             | DELETE |
| get all sub-categories                    | /sub-categories                               | GET    |
| get sub-categories belong to the category | /sub-categories?category-name={category name} | GET    |
| add sub-category                          | /sub-categories/:slug                         | POST   |
| update sub-category                       | /sub-categories/:slug                         | PUT    |
| delete sub-category                       | /sub-categories/:slug                         | GET    |
| get all posts                             | /posts                                        | GET    |
| get posts belong to the category          | /posts?category-name={category name}          | GET    |
| get posts belongs to the sub-category     | /posts?sub-category-name={sub-category name}  | GET    |
| add post                                  | /posts/:slug                                  | POST   |
| update post                               | /posts/:slug                                  | PUT    |
| delete post                               | /post/:slug                                   | DELETE |

