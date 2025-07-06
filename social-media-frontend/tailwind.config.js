module.exports = {
  theme: {
    extend: {
      colors: {
        primary: '#1da1f2', // Twitter-like blue
        secondary: '#657786', // Dark gray for secondary elements
        accent: '#e1e8ed', // Light gray
        background: '#f7f9fa', // Soft background color
        danger: '#e0245e', // Red for errors or notifications
      },
      fontFamily: {
        sans: ['Inter', 'Helvetica', 'Arial', 'sans-serif'],
      },
      spacing: {
        18: '4.5rem',
        22: '5.5rem',
        28: '7rem',
      },
      screens: {
        xs: '400px', // Small mobile devices
        sm: '640px', // Small screens (tablets)
        md: '768px', // Medium screens (laptops)
        lg: '1024px', // Large screens (desktops)
        xl: '1280px', // Extra large screens
        '2xl': '1536px',
      },
    },
  },
  plugins: [],
}
