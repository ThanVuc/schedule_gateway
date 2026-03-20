# Team API - Structured For FE

Base URL: /api/v1

Enum quick map:
- GroupRole: 0 UNSPECIFIED, 1 OWNER, 2 MANAGER, 3 MEMBER, 4 VIEWER
- SprintStatus: 0 UNSPECIFIED, 1 DRAFT, 2 ACTIVE, 3 COMPLETED, 4 CANCELLED
- WorkStatus: 0 UNSPECIFIED, 1 TODO, 2 IN_PROGRESS, 3 IN_REVIEW, 4 DONE
- WorkPriority: 0 UNSPECIFIED, 1 LOW, 2 MEDIUM, 3 HIGH

## 1. API 1 - Ping() - Endpoints: GET /api/v1/ts/groups/ping

- Request:
  - No body
- Response (200):

```json
{
  "message": "pong"
}
```

## 2. API 2 - CreateGroup() - Endpoints: POST /api/v1/ts/groups

- Request body:

```json
{
  "name": "Platform Team",
  "description": "Owns core platform"
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Group created successfully",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{group_id}}",
      "name": "Platform Team",
      "description": "Owns core platform",
      "owner": {
        "id": "{{owner_id}}",
        "email": "owner@example.com",
        "avatar": ""
      },
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T10:00:00Z"
    }
  }
}
```

## 3. API 3 - GetGroup() - Endpoints: GET /api/v1/ts/groups/:group_id

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "Group retrieved successfully",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{group_id}}",
      "name": "Platform Team",
      "description": "Owns core platform",
      "owner": {
        "id": "{{owner_id}}",
        "email": "owner@example.com",
        "avatar": ""
      },
      "my_role": 1,
      "active_sprint": "",
      "avatar": "",
      "members_total": 3,
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T10:00:00Z"
    }
  }
}
```

## 4. API 4 - UpdateGroup() - Endpoints: PATCH /api/v1/ts/groups/:group_id

- Request body:

```json
{
  "name": "Platform Team Updated",
  "description": "Updated description"
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Group updated successfully",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{group_id}}",
      "name": "Platform Team Updated",
      "description": "Updated description",
      "owner": {
        "id": "{{owner_id}}",
        "email": "owner@example.com",
        "avatar": ""
      },
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T11:00:00Z"
    }
  }
}
```

## 5. API 5 - DeleteGroup() - Endpoints: DELETE /api/v1/ts/groups/:group_id

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "Group deleted successfully",
  "reasonStatusCode": "success",
  "metadata": {
    "item": null
  }
}
```

## 6. API 6 - ListMembers() - Endpoints: GET /api/v1/ts/groups/:group_id/members

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "Group members retrieved successfully",
  "reasonStatusCode": "success",
  "metadata": {
    "items": [
      {
        "id": "{{user_id}}",
        "email": "member@example.com",
        "avatar": "",
        "role": 3,
        "joined_at": "2026-03-20T10:00:00Z"
      }
    ],
    "total": 1
  }
}
```

## 7. API 7 - UpdateMemberRole() - Endpoints: PATCH /api/v1/ts/groups/:group_id/members/:user_id

- Request body:

```json
{
  "role": 2
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Member role updated successfully",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{user_id}}",
      "email": "member@example.com",
      "avatar": "",
      "role": 2,
      "joined_at": "2026-03-20T10:00:00Z"
    }
  }
}
```

## 8. API 8 - CreateSprint() - Endpoints: POST /api/v1/ts/groups/:group_id/sprints

- Request body:

```json
{
  "name": "Sprint 12",
  "goal": "Release milestone",
  "start_date": "2026-04-01",
  "end_date": "2026-04-14"
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Create sprint successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{sprint_id}}",
      "group_id": "{{group_id}}",
      "name": "Sprint 12",
      "goal": "Release milestone",
      "status": 1,
      "start_date": "2026-04-01",
      "end_date": "2026-04-14",
      "total_work": 0,
      "completed_work": 0,
      "progress_percent": 0,
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T10:00:00Z"
    }
  }
}
```

## 9. API 9 - ListSprints() - Endpoints: GET /api/v1/ts/groups/:group_id/sprints

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "List sprints successful",
  "reasonStatusCode": "success",
  "metadata": {
    "items": [
      {
        "id": "{{sprint_id}}",
        "group_id": "{{group_id}}",
        "name": "Sprint 12",
        "goal": "Release milestone",
        "status": 1,
        "start_date": "2026-04-01",
        "end_date": "2026-04-14",
        "total_work": 0,
        "completed_work": 0,
        "progress_percent": 0,
        "created_at": "2026-03-20T10:00:00Z",
        "updated_at": "2026-03-20T10:00:00Z"
      }
    ],
    "total": 1
  }
}
```

## 10. API 10 - GetSprint() - Endpoints: GET /api/v1/ts/groups/:group_id/sprints/:sprint_id

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "Get sprint successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{sprint_id}}",
      "group_id": "{{group_id}}",
      "name": "Sprint 12",
      "goal": "Release milestone",
      "status": 1,
      "start_date": "2026-04-01",
      "end_date": "2026-04-14",
      "total_work": 0,
      "completed_work": 0,
      "progress_percent": 0,
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T10:00:00Z"
    }
  }
}
```

## 11. API 11 - UpdateSprint() - Endpoints: PATCH /api/v1/ts/groups/:group_id/sprints/:sprint_id

- Request body:

```json
{
  "name": "Sprint 12.1",
  "goal": "Update scope"
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Update sprint successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{sprint_id}}",
      "group_id": "{{group_id}}",
      "name": "Sprint 12.1",
      "goal": "Update scope",
      "status": 1,
      "start_date": "2026-04-01",
      "end_date": "2026-04-14",
      "total_work": 0,
      "completed_work": 0,
      "progress_percent": 0,
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T11:00:00Z"
    }
  }
}
```

## 12. API 12 - UpdateSprintStatus() - Endpoints: PATCH /api/v1/ts/groups/:group_id/sprints/:sprint_id/status

- Request body:

```json
{
  "status": 2
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Update sprint status successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{sprint_id}}",
      "status": 2
    }
  }
}
```

## 13. API 13 - DeleteSprint() - Endpoints: DELETE /api/v1/ts/groups/:group_id/sprints/:sprint_id

- Request:
  - No body
- Response (204):

```json
{
  "statusCode": 204,
  "message": "Delete sprint successful",
  "reasonStatusCode": "no content",
  "metadata": {
    "item": {
      "is_success": true
    }
  }
}
```

## 14. API 14 - CreateWork() - Endpoints: POST /api/v1/ts/groups/:group_id/works

- Request body:

```json
{
  "name": "Implement Team API",
  "description": "Build endpoints and docs",
  "sprint_id": "{{sprint_id}}"
}
```

- Response (201):

```json
{
  "statusCode": 201,
  "message": "Create work successful",
  "reasonStatusCode": "created",
  "metadata": {
    "item": {
      "id": "{{work_id}}",
      "name": "Implement Team API",
      "description": "Build endpoints and docs",
      "status": 1,
      "sprint": {
        "id": "{{sprint_id}}",
        "name": "Sprint 12"
      },
      "assignee": {},
      "story_point": "",
      "due_date": "",
      "check_list": {
        "total": 0,
        "completed": 0,
        "items": []
      },
      "comments": {
        "total": 0,
        "comments": []
      },
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T10:00:00Z",
      "version": 0
    }
  }
}
```

## 15. API 15 - ListWorks() - Endpoints: GET /api/v1/ts/groups/:group_id/works?sprint_id=:sprint_id

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "List works successful",
  "reasonStatusCode": "success",
  "metadata": {
    "items": [
      {
        "id": "{{work_id}}",
        "name": "Implement Team API",
        "description": "Build endpoints and docs",
        "status": 1,
        "sprint": {
          "id": "{{sprint_id}}",
          "name": "Sprint 12"
        },
        "assignee": {},
        "story_point": "",
        "due_date": "",
        "check_list": {
          "total": 0,
          "completed": 0,
          "items": []
        },
        "comments": {
          "total": 0,
          "comments": []
        },
        "created_at": "2026-03-20T10:00:00Z",
        "updated_at": "2026-03-20T10:00:00Z",
        "version": 0
      }
    ],
    "total": 1
  }
}
```

## 16. API 16 - GetWork() - Endpoints: GET /api/v1/ts/groups/:group_id/works/:work_id

- Request:
  - No body
- Response (200):

```json
{
  "statusCode": 200,
  "message": "Get work successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{work_id}}",
      "name": "Implement Team API",
      "description": "Build endpoints and docs",
      "status": 1,
      "sprint": {
        "id": "{{sprint_id}}",
        "name": "Sprint 12"
      },
      "assignee": {},
      "story_point": "",
      "due_date": "",
      "check_list": {
        "total": 0,
        "completed": 0,
        "items": []
      },
      "comments": {
        "total": 0,
        "comments": []
      },
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T10:00:00Z",
      "version": 0
    }
  }
}
```

## 17. API 17 - UpdateWork() - Endpoints: PATCH /api/v1/ts/groups/:group_id/works/:work_id

- Request body:

```json
{
  "name": "Implement Team API v2",
  "description": "Improve validation and docs",
  "assignee_id": "{{assignee_id}}",
  "status": 2,
  "story_point": "8",
  "due_date": "2026-04-20",
  "priority": 3,
  "version": 1
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Update work successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{work_id}}",
      "name": "Implement Team API v2",
      "description": "Improve validation and docs",
      "status": 2,
      "sprint": {
        "id": "{{sprint_id}}",
        "name": "Sprint 12"
      },
      "assignee": {
        "id": "{{assignee_id}}",
        "email": "assignee@example.com",
        "avatar": ""
      },
      "story_point": "8",
      "due_date": "2026-04-20",
      "check_list": {
        "total": 0,
        "completed": 0,
        "items": []
      },
      "comments": {
        "total": 0,
        "comments": []
      },
      "created_at": "2026-03-20T10:00:00Z",
      "updated_at": "2026-03-20T11:00:00Z",
      "version": 2
    }
  }
}
```

## 18. API 18 - DeleteWork() - Endpoints: DELETE /api/v1/ts/groups/:group_id/works/:work_id

- Request:
  - No body
- Response (204):

```json
{
  "statusCode": 204,
  "message": "Delete work successful",
  "reasonStatusCode": "no content",
  "metadata": {
    "item": {
      "is_success": true
    }
  }
}
```

## 19. API 19 - CreateChecklistItem() - Endpoints: POST /api/v1/ts/groups/:group_id/works/:work_id/checklists

- Request body:

```json
{
  "name": "Write API tests"
}
```

- Response (201):

```json
{
  "statusCode": 201,
  "message": "Create checklist item successful",
  "reasonStatusCode": "created",
  "metadata": {
    "item": {
      "id": "{{checklist_id}}",
      "name": "Write API tests",
      "is_completed": false
    }
  }
}
```

## 20. API 20 - UpdateChecklistItem() - Endpoints: PATCH /api/v1/ts/groups/:group_id/works/:work_id/checklists/:checklist_id

- Request body:

```json
{
  "name": "Write integration tests",
  "is_completed": true
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Update checklist item successful",
  "reasonStatusCode": "success",
  "metadata": {
    "item": {
      "id": "{{checklist_id}}",
      "name": "Write integration tests",
      "is_completed": true
    }
  }
}
```

## 21. API 21 - DeleteChecklistItem() - Endpoints: DELETE /api/v1/ts/groups/:group_id/works/:work_id/checklists/:checklist_id

- Request:
  - No body
- Response (204):

```json
{
  "statusCode": 204,
  "message": "Delete checklist item successful",
  "reasonStatusCode": "no content",
  "metadata": {
    "item": {
      "id": "{{checklist_id}}",
      "name": "Write integration tests",
      "is_completed": true
    }
  }
}
```

## 22. API 22 - CreateComment() - Endpoints: POST /api/v1/ts/groups/:group_id/works/:work_id/comments

- Request body:

```json
{
  "content": "Please review this task"
}
```

- Response (201):

```json
{
  "statusCode": 201,
  "message": "Create comment successful",
  "reasonStatusCode": "created",
  "metadata": {
    "items": [
      {
        "id": "{{comment_id}}",
        "content": "Please review this task",
        "creator": {
          "id": "{{creator_id}}",
          "email": "creator@example.com",
          "avatar": ""
        },
        "created_at": "2026-03-20T10:00:00Z"
      }
    ],
    "total": 1
  }
}
```

## 23. API 23 - UpdateComment() - Endpoints: PATCH /api/v1/ts/groups/:group_id/works/:work_id/comments/:comment_id

- Request body:

```json
{
  "content": "Updated comment body"
}
```

- Response (200):

```json
{
  "statusCode": 200,
  "message": "Update comment successful",
  "reasonStatusCode": "success",
  "metadata": {
    "items": [
      {
        "id": "{{comment_id}}",
        "content": "Updated comment body",
        "creator": {
          "id": "{{creator_id}}",
          "email": "creator@example.com",
          "avatar": ""
        },
        "created_at": "2026-03-20T10:00:00Z"
      }
    ],
    "total": 1
  }
}
```

## 24. API 24 - DeleteComment() - Endpoints: DELETE /api/v1/ts/groups/:group_id/works/:work_id/comments/:comment_id

- Request:
  - No body
- Response (204):

```json
{
  "statusCode": 204,
  "message": "Delete comment successful",
  "reasonStatusCode": "no content",
  "metadata": {
    "items": [],
    "total": 0
  }
}
```

## Notes

- You asked to skip headers, so only request body is shown for APIs that send body.
- GET and DELETE APIs are listed with "No body".
- Error format can be either standard wrapper or {"error": "..."} for some Group controller local errors.
