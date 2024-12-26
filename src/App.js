import React, { useState } from "react";
import "./App.css"; // Include this for custom styles

function App() {
  const [formData, setFormData] = useState({
    cabinNumber: "",
    firstName: "",
    lastName: "",
    dateOfBirth: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Form Data Submitted: ", formData);
    alert("Form submitted successfully!");
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Holland Cruise Lines</h1>
        <nav>
          <ul>
            <li>Home</li>
            <li>Destinations</li>
            <li>Contact</li>
          </ul>
        </nav>
      </header>

      <main>
        <h2>Guest Information</h2>
        <form className="submission-form" onSubmit={handleSubmit}>
          <label>
            Cabin Number:
            <input
              type="text"
              name="cabinNumber"
              value={formData.cabinNumber}
              onChange={handleChange}
              required
            />
          </label>

          <label>
            First Name:
            <input
              type="text"
              name="firstName"
              value={formData.firstName}
              onChange={handleChange}
              required
            />
          </label>

          <label>
            Last Name:
            <input
              type="text"
              name="lastName"
              value={formData.lastName}
              onChange={handleChange}
              required
            />
          </label>

          <label>
            Date of Birth:
            <input
              type="date"
              name="dateOfBirth"
              value={formData.dateOfBirth}
              onChange={handleChange}
              required
            />
          </label>

          <button type="submit">Submit</button>
        </form>
      </main>
    </div>
  );
}

export default App;
