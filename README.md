# Recruitment-Management-System
<ul>
<li><h3><u>Task:</u> <i>Create a backend server for Recruitment Management System</i></h3>
</li>
</ul>

<h2>
Requirements:
</h2>
<ol>
<li>
Users can create their profile and upload resumes (only PDF and DOCX formats allowed)
</li>
<li>
Uploaded resumes will be processed using a third-party API, and relavent information will be stored in the database.
</li>
<li>
Admin users can create job openings.
</li>
<li>
Admin users can view all uploaded resumes and extracted data of applicants.
</li>
<li>
Applicants can view job openings.
</li>
<li>
Applicants can view job openings.
</li>
</ol>

<h2>
APIs
</h2>

<ol>
<li>
<b>POST /signup:</b> Create a profile on the system (Name, Email, Password, UserType
(Admin/Applicant), Profile Headline, Address).
</li>
<li>
<b>POST /login:</b> Authenticate users and return a JWT token upon successful validation.
</li>
<li>
<b>POST /uploadResume:</b> Authenticated API for uploading resume files (only PDF or DOCX) of
the applicant. Only Applicant type users can access this API.</li>
<li>
<b>POST /admin/job:</b> Authenticated API for creating job openings. Only Admin type users can
access this API.
</li>
<li>
<b>GET /admin/job/{job_id}:</b> Authenticated API for fetching information regarding a job
opening. Returns details about the job opening and a list of applicants. Only Admin type
users can access this API.
</li>
<li>
<b>GET /admin/applicants:</b> Authenticated API for fetching a list of all users in the system. Only
Admin type users can access this API.
</li>
<li>
<b>GET /admin/applicant/{applicant_id}:</b> Authenticated API for fetching extracted data of an
applicant. Only Admin type users can access this API.
</li>
<li><b>GET /jobs:</b> Authenticated API for fetching job openings. All users can access this API.
</li>
<li><b>GET /jobs/apply?job_id={job_id}:</b> Authenticated API for applying to a particular job. Only
Applicant users are allowed to apply for jobs.
</li>
</ol>

<h2>Models:</h2>
<ul>
    <li>
    User:
        <ul>
        <li>Name: string</li>
        <li>Email: string</li>
        <li>Address: string</li>
        <li>UserType (Applicant/Admin)</li>
        <li>PasswordHash: string</li>
        <li>Profile Headline: string</li>
        <li>Profile: Profile</li>
        </ul>
    </li>
    <li>
    Profile:
        <ul>
        <li>Applicant: User</li>
        <li>Resume File Address: string</li>
        <li>Skills: string</li>
        <li>Education: string</li>
        <li>Experience: string</li>
        <li>Name: string</li>
        <li>Email: string</li>
        <li>Phone: string</li>
        </ul>
    <li>
    Job:
        <ul>
        <li>Title: string</li>
        <li>Description: string</li>
        <li>PostedOn: datetime</li>
        <li>TotalApplications: int</li>
        <li>Company Name: string</li>
        <li>PostedBy: User</li>
        </ul>
    </li>
</ul>

<h2>Usage instructions:</h2> Send a request to the API with the resume file (PDF or DOCX), it will return a
JSON with all the scrapped information as shown below:

```
{
"education": [
    {
    "name": "Wharton School of the University of Pennsylvania",
    "url": "http://dbpedia.org/page/Wharton_School_of_the_University_of_Pennsylvania"
    },

    {
    "name": "Penn's College of Arts and Sciences",
    "url": "http://dbpedia.org/page/University_of_Pennsylvania_School_of_Arts_and_Sciences"
    }
],

"email": "elonmusk@teslamotors.com",

"experience": [
{
    "dates": [
        "2006-06"
    ],
    "name": "Solar City"
},
{
    "dates": [
        "06-2002"
    ],
    "name": "SpaceX",
    "url": "http://spacex.com"
},
{
"dates": [
    "03-1999",
    "10-2002"
    ],
"name": "X.com and Paypal",
"url": "http://paypal.com"
}
],
    "name": "Elon Musk",
    "phone": "650 68100",
    "skills": [
    "Entrepreneurship",
    "Innovation",
    "Mars",
    "Space",
    "Electric Cars",
    "Physics",
    "Maths",
    "Calculus",
    "Distrupting Technologies"
    ]
}
```