from typing import Optional

import openai
from openai import OpenAI
from dotenv import load_dotenv
import os
import json
from pydantic import BaseModel, HttpUrl, constr, conint

# Enums for EmploymentType and JobType
from enum import Enum
from sqlalchemy import create_engine, Column, String, Integer, Enum as SqlEnum, DateTime, Boolean
from sqlalchemy.orm import sessionmaker

from crawler.models.base import Base

# Load environment variables from .env file
load_dotenv()

# Set your OpenAI API key from environment variable
openai.api_key = os.getenv('OPENAI_API_KEY')
client = OpenAI()


class EmploymentType(str, Enum):
    full_time = "Full-time"
    part_time = "Part-time"
    contract = "Contract"
    internship = "Internship"
    freelance = "Freelance"


class JobType(str, Enum):
    remote = "Remote"
    on_site = "On-site"
    hybrid = "Hybrid"



class JobPost(BaseModel):
    jobId: str
    jobTitle: str
    employmentType: EmploymentType  # Enum for employment type
    salaryRange: str
    jobType: JobType
    contractLength: Optional[str] = None
    industry: Optional[str] = None
    requiredSkills: Optional[str] = None
    educationLevel: Optional[str] = None
    jobPostedDate: str
    jobDescription: str
    applicationLink: str  # URL for applying to the job, up to 200 characters
    isRequireUsAPerson: bool  # Changed to bool for better type accuracy
    isSecurityClearanceRequired: bool  # Changed to bool for better type accuracy
    visaSponsorshipDetails: Optional[str] = None
    minimumExperienceRequired: int
    applicants: int
    hirer: str
    hirerProfileLink: Optional[str] = None
    location: str
    atsKeywords: Optional[str] = None

# Define the database URL
DATABASE_URL = f"postgresql://{os.getenv('DB_USERNAME')}:{os.getenv('DB_PASSWORD')}@{os.getenv('DB_HOST')}:{os.getenv('DB_PORT')}/{os.getenv('DB_NAME')}"

# Create the SQLAlchemy engine
engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

if __name__ == "__main__":
    completion = client.beta.chat.completions.parse(
        model="gpt-4o-mini",
        messages=[
            {
                "role": "system",
                "content": "The current time is: 2024-10-01T01:22:00Z. When processing job-related data, please calculate the 'jobPostedDate' based on the current time and the phrases like 'posted X hours ago'. Your task is to process unstructured job-related data scraped from various websites and convert it into a structured JSON format suitable for populating a database. The JSON format must adhere to the 'JobPost' table schema outlined below. Each field must comply with the specified constraints, and missing or unavailable data should be represented as empty strings or null values. Ensure the output is valid JSON format.\n\nSchema for 'JobPost':\n2. 'jobId': (string, max 50 chars) Unique Job ID for each post.\n\n3. 'jobTitle': (string, max 50 chars) Title of the job.\n\n4. 'employmentType': (enum) Employment type: 'Full-time', 'Part-time', 'Contract', or 'Internship'.\n\n5. 'salaryRange': (string, max 50 chars) Salary or hourly rate range.\n\n6. 'jobType': (enum) Job type: 'On-site', 'Remote', or 'Hybrid'.\n\n7. 'contractLength': (string, max 50 chars) Contract duration for contract jobs.\n\n8. 'industry': (string, max 50 chars) Relevant industry (e.g., IT, Healthcare).\n\n9. 'requiredSkills': (string, max 200 chars) Skills required for the job.\n\n10. 'educationLevel': (string, max 50 chars) Minimum education level required.\n\n11. 'jobPostedDate': (datetime)  estimated from text (e.g., 'posted X hours ago', 'posted yesterday', 'posted on October 1, 2023') or Date when the job was posted. Be sure to parse variations in phrasing correctly to derive the exact date.\n\n12. 'jobDescription': (string, max 1000 chars) Detailed job description.\n\n13. 'applicationLink': (string, max 200 chars) URL link to the job application.\n\n14. 'isRequireUSAPerson': (boolean) Whether U.S. citizenship or permanent residency is required.\n\n15. 'isSecurityClearanceRequired': (boolean) Whether security clearance is required.\n\n16. 'visaSponsorshipDetails': (string, max 50 chars) Visa sponsorship details, if applicable.\n\n17. 'minimumExperienceRequired': (integer) Minimum experience required (in years).\n\n18. 'applicants': (integer) Number of applicants so far.\n\n19. 'hirer': (string, max 30 chars) Name of the hiring person/company.\n\n20. 'hirerProfileLink': (string, max 100 chars) Profile link of the hirer (e.g., LinkedIn or company page).\n\n21. 'location': (string, max 30 chars) Job location.\n\n22. 'atsKeywords': (string, max 200 chars) extract relevant keywords that can enhance a resume's visibility in Applicant Tracking Systems (ATS). Focus on the following types of keywords:\n\n1. **Job Title:** Extract the exact job title and common variations.\n2. **Technical Skills:** Identify programming languages, frameworks, tools, and platforms mentioned in the job description.\n3. **Soft Skills:** Extract any interpersonal skills or attributes that are emphasized.\n4. **Industry-Specific Terms:** Look for jargon or terms that are specific to the industry of the job posting.\n5. **Certifications:** Note any required or preferred certifications related to the position.\n6. **Achievements and Metrics:** Extract quantitative results or accomplishments mentioned in the job description.\n7. **Action Verbs:** Identify strong action verbs that convey responsibilities and contributions.\n\nOutput the extracted keywords in a structured format, ensuring they are relevant to the job title and aligned with the skills and qualifications sought in the posting. Remove duplicates and organize keywords into categories (e.g., Technical Skills, Soft Skills, Certifications)..\n\nOutput each job in a structured JSON format matching these fields. Validate the data for compliance with the specified constraints. Make sure all data values's are below max chars limits and return Only JSON formated data to directly feed in the system."
            },
            {
                "role": "user",
                "content": "{'company_name': 'Gravity IT Resources', 'primary_description': 'Texas, United States · 7 hours ago · Over 100 applicants', 'hirer_name': None, 'linkedin_profile_link': None, 'job_insight': '$45/hr - $55/hr  Remote\nMatches your job preferences, workplace type is Remote.\n  Contract\nMatches your job preferences, job type is Contract.\n  Mid-Senior level', 'job_description': 'About the job\n\nJob Title: QA Engineer\n\nLocation: Remote in Texas\n\nJob Type: Contract\n\n\n\n\nEmployment Eligibility: Gravity cannot transfer nor sponsor a work visa for this position. Applicants must be eligible to work in the U.S. for any employer directly \n\n\n\n\nRole & Responsibilities:\n\nPerform software quality assurance, testing, and support for a wide range of social services applications and other related platforms.\nDevelop, enhance, maintain, and execute test plans using Agile methodologies on the SOA (Service-Oriented Architecture) platform.\nParticipate in various phases of the SDLC, including requirements analysis, test design, test data preparation, and actual testing.\nCollaborate with other teams and create technical documentation.\nProvide technical support during testing and implementation phases.\nEnsure that applications meet technical standards and function properly.\nAttend meetings to give input on program design, coding, and testing.\n\n\n\n\nRequired Experience and Skills\n\n8+ years of experience in software development and SDLC.\nDeep understanding of software engineering practices, including Agile methodologies.\nExpertise in analyzing requirements and creating test plans to ensure the applications meet their intended functionality.\nExperience with automated testing tools such as HP Unified Functional Testing (UFT), Selenium, SoapUI, and SQL.\nExperience working in an Agile environment.\nStrong communication skills (both written and oral).\n\n\n\n\nPreferred Skills\n\n3+ years of experience with specific tools like HP LoadRunner, Katalon, Selenium, JIRA, SQL, JAMA, and manual or automated testing.\nExperience working with various teams in a collaborative environment.', 'company_description': 'Gravity IT Resources provides the consulting expertise and IT talent that powers digital transformation. We employ consistent, innovative strategies across multiple practice areas to help our clients leverage technology to drive bottom line results.\n\nWe offer contract/staff augmentation, contract-to-hire and direct hire/direct placement staffing services of IT professionals. These services are tailored to meet the specific needs of several different market segments including Commercial, Healthcare and Government.\n\nWe have specialized teams and resources for delivering our staffing services for specific IT skill sets including Salesforce, ServiceNow and Digital Transformation.\n\nWe also have a Cloud Analytics consulting practice focused on helping clients leverage modern cloud data and analytics tools to improve corporate decision making in Marketing, Sales and Operations functions.\n…\nshow more', 'applicants': 'Applicants for this job\n1,373\tApplicants\n1,373\tApplicants in the past day'}"
            },
        ],
        response_format=JobPost
    )




    content = completion.choices[0].message.content  # Accessing the content correctly

    # Replace single quotes with double quotes for JSON compatibility
    content = content.replace("'", "\"")

    # Parse the JSON content
    job_data = json.loads(content)

    # Create an instance of JobPost
    job_post = JobPost(
        jobId=job_data.get('jobId', ''),
        jobTitle=job_data.get('jobTitle', ''),
        employmentType=EmploymentType(job_data.get('employmentType', '')),
        salaryRange=job_data.get('salaryRange', ''),
        jobType=JobType(job_data.get('jobType', '')),
        contractLength=job_data.get('contractLength', ''),
        industry=job_data.get('industry', ''),
        requiredSkills=job_data.get('requiredSkills', ''),
        educationLevel=job_data.get('educationLevel', ''),
        jobPostedDate=job_data.get('jobPostedDate', ''),
        jobDescription=job_data.get('jobDescription', ''),
        applicationLink=job_data.get('applicationLink', ''),
        isRequireUsAPerson=job_data.get('isRequireUsAPerson', False),
        isSecurityClearanceRequired=job_data.get('isSecurityClearanceRequired', False),
        visaSponsorshipDetails=job_data.get('visaSponsorshipDetails', ''),
        minimumExperienceRequired=job_data.get('minimumExperienceRequired', 0),
        applicants=job_data.get('applicants', 0),
        hirer=job_data.get('hirer', ''),
        hirerProfileLink=job_data.get('hirerProfileLink', ''),
        location=job_data.get('location', ''),
        atsKeywords=job_data.get('atsKeywords', '')
    )

    print(job_post)