# Gono Frontend - Feature-Sliced Design Architecture

## 📁 Cấu trúc dự án

Dự án đã được refactor theo kiến trúc **Feature-Sliced Design (FSD)** - một phương pháp tổ chức code frontend hiện đại, scalable và maintainable.

```
src/
├── app/                    # 🚀 Application Layer
│   ├── main.js            # Entry point
│   ├── App.vue            # Root component
│   ├── router/            # Routing configuration
│   ├── config/            # App-level configuration
│   └── styles/            # Global styles & assets
│
├── pages/                  # 📄 Pages Layer
│   ├── feed/              # Feed page
│   ├── login/             # Login page
│   ├── register/          # Register page
│   ├── profile/           # Profile page
│   ├── submit/            # Submit post page
│   ├── post-detail/       # Post detail page
│   ├── teams/             # Teams page
│   └── notifications/     # Notifications page
│
├── widgets/                # 🧩 Widgets Layer
│   ├── post-card/         # Post card component
│   ├── sidebar/           # Sidebar navigation
│   ├── header/            # Header component
│   └── trending-topics/   # Trending topics widget
│
├── features/               # ⚡ Features Layer
│   ├── auth/              # Authentication feature
│   │   ├── api/          # Auth API calls
│   │   ├── model/        # Auth store & logic
│   │   └── ui/           # Auth UI components (LoginForm, RegisterForm)
│   ├── create-post/       # Create post feature
│   ├── upvote-post/       # Upvote functionality
│   └── edit-profile/      # Edit profile feature
│
├── entities/               # 🎯 Entities Layer
│   ├── user/              # User entity
│   │   ├── api/          # User API
│   │   ├── model/        # User store
│   │   └── ui/           # User UI components (UserAvatar)
│   ├── post/              # Post entity
│   │   ├── api/          # Post API
│   │   ├── model/        # Post store
│   │   └── ui/           # Post UI components
│   └── team/              # Team entity
│       ├── api/
│       ├── model/
│       └── ui/
│
└── shared/                 # 🔧 Shared Layer
    ├── api/               # API client & configuration
    │   └── client.js     # Axios instance with interceptors
    ├── ui/                # Reusable UI components
    ├── lib/               # Utility functions
    │   └── utils.js      # Common utilities (formatters, validators)
    └── config/            # Shared configuration
```

## 🎯 Nguyên tắc FSD

### 1. **Import Rule (từ dưới lên trên)**
```javascript
// ✅ ĐÚNG: Layer trên có thể import layer dưới
pages/ → widgets/
widgets/ → features/
features/ → entities/
entities/ → shared/

// ❌ SAI: Không được import ngược lại
shared/ → entities/ ❌
entities/ → features/ ❌
```

### 2. **Separation of Concerns**
- **app/**: Khởi tạo app, config toàn cục
- **pages/**: Các routes/pages, kết nối widgets
- **widgets/**: Composite UI blocks, kết hợp nhiều features
- **features/**: User scenarios (login, create post, upvote)
- **entities/**: Business entities (user, post, team)
- **shared/**: Code dùng chung (utils, API, UI components)

### 3. **Public API**
Mỗi slice nên export public API qua index.js:

```javascript
// entities/user/index.js
export { useUserStore } from './model/userStore'
export { userApi } from './api/userApi'
export { default as UserAvatar } from './ui/UserAvatar.vue'
```

## 📝 Ví dụ Flow: Login

### 1. **User truy cập `/login`**
```
pages/login/ui/LoginPage.vue
```

### 2. **LoginPage render LoginForm**
```
features/auth/ui/LoginForm.vue
```

### 3. **User submit form → gọi authStore**
```
features/auth/model/authStore.js → login()
```

### 4. **authStore gọi API**
```
features/auth/api/authApi.js → login()
shared/api/client.js (axios instance)
```

### 5. **API success → Update entities**
```
authStore.setAccessToken()
userStore.setUser() (entities/user/model)
```

### 6. **Router guard redirect**
```
app/router/index.js → beforeEach → /feed
```

## 🔥 Các Store chính

### **authStore** (features/auth/model)
- Quản lý authentication state
- Lưu trữ access token
- Login/logout logic

### **userStore** (entities/user/model)
- Quản lý user entity
- User profile data
- User state (isAuthenticated, isAdmin)

### **postStore** (entities/post/model)
- Quản lý posts
- CRUD operations
- Post state (loading, error)

## 🛠️ Utilities (shared/lib/utils.js)

```javascript
// Email validation
validateEmail(email)

// Password validation  
validatePassword(password)

// Get user initials
getUserInitials(name)

// Format timestamp
formatTimestamp(timestamp)

// Truncate text
truncateText(text, maxLength)
```

## 🔌 API Configuration

### Setup Interceptors
```javascript
// app/config/api.js
import { setupAuthInterceptor, setupResponseInterceptor } from '@/shared/api/client'

setupAuthInterceptor(() => authStore.token)
setupResponseInterceptor(() => {
  authStore.clearAuth()
  window.location.href = '/login'
})
```

## 🚀 Migration từ cấu trúc cũ

### Cũ (Feature-based)
```
src/features/auth/
  ├── api/authService.js
  ├── stores/authStore.js
  ├── composables/useAuth.js
  └── views/Login.vue
```

### Mới (FSD)
```
src/
├── features/auth/
│   ├── api/authApi.js
│   ├── model/authStore.js
│   └── ui/LoginForm.vue
├── entities/user/
│   ├── model/userStore.js
│   └── ui/UserAvatar.vue
└── pages/login/
    └── ui/LoginPage.vue
```

## 📚 Tài liệu tham khảo

- [Feature-Sliced Design](https://feature-sliced.design/)
- [FSD Examples](https://github.com/feature-sliced/examples)

## ✨ Lợi ích của FSD

1. **Scalability**: Dễ dàng mở rộng khi dự án lớn
2. **Maintainability**: Code được tổ chức rõ ràng, dễ bảo trì
3. **Reusability**: Tái sử dụng code hiệu quả
4. **Testability**: Dễ dàng test từng layer riêng biệt
5. **Team collaboration**: Nhiều người có thể làm việc song song không conflict
6. **No circular dependencies**: Ngăn chặn circular imports
