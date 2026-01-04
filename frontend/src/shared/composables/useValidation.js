export function useValidation() {
  const validateEmail = (email) => {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    return re.test(email)
  }

  const validatePassword = (password) => {
    return password.length >= 8
  }

  const validateUsername = (username) => {
    return username.length >= 3
  }

  const validateRegisterForm = (form) => {
    const errors = {}

    if (!form.username || !validateUsername(form.username)) {
      errors.username = 'Username must be at least 3 characters'
    }

    if (!form.email) {
      errors.email = 'Email is required'
    } else if (!validateEmail(form.email)) {
      errors.email = 'Please enter a valid email address'
    }

    if (!form.password) {
      errors.password = 'Password is required'
    } else if (!validatePassword(form.password)) {
      errors.password = 'Password must be at least 8 characters'
    }

    if (form.password !== form.confirm_password) {
      errors.confirm_password = 'Passwords do not match'
    }

    return errors
  }

  const validateLoginForm = (form) => {
    const errors = {}

    if (!form.email) {
      errors.email = 'Email is required'
    } else if (!validateEmail(form.email)) {
      errors.email = 'Please enter a valid email address'
    }

    if (!form.password) {
      errors.password = 'Password is required'
    }

    return errors
  }

  return {
    validateEmail,
    validatePassword,
    validateUsername,
    validateRegisterForm,
    validateLoginForm
  }
}
