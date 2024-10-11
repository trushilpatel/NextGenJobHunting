from typing import Optional

import openai
from openai import OpenAI
from dotenv import load_dotenv
import os
import json
from pydantic import BaseModel

# Enums for EmploymentType and JobType
from enum import Enum
from sqlalchemy import create_engine
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
                "content": """{
  "job_id": 4048357252,
  "job_data": {
    "company_name": "Adtalem Global Education",
    "primary_description": "Minneapolis, MN · 34 minutes ago · 0 applicants",
    "hirer_name": null,
    "linkedin_profile_link": null,
    "job_insight": "$70.7K/yr - $124.2K/yr  On-site  Full-time  Associate",
    "job_description": "About the job\n\nCompany Description\n\nAbout Chamberlain University\n\nChamberlain University has a 130-year history and the mission to educate, empower, and embolden diverse healthcare professionals who advance the health of people, families, communities, and nations. Our regionally accredited University is comprised of the College of Nursing and the College of Health Professions. Our College of Nursing offers a Bachelor of Science in Nursing (BSN) degree program at 23 locations and online post-licensure degree programs and certificates including RN to BSN, RN to MSN, a Master of Science in Nursing (MSN) with six Specialty Tracks, and a Doctor of Nursing Practice (DNP). Our College of Health Professions offers online options with a Master of Public Health degree, Master of Physician Assistant Studies, Public Health Certificate programs, and a Master of Social Work degree. Learn more about Chamberlain University, part of Adtalem Global Education, at www.chamberlain.edu.\n\nThere is one very significant way Chamberlain is different than other Universities - Chamberlain Care®. Healthcare education is offered in a culture of service excellence and care in four ways: care for self, care for colleagues, care for patients/clients, care for all. By living Chamberlain Care®, we’ve built a culture in which colleagues and students thrive and we graduate extraordinary healthcare professionals who transform the health of people worldwide.\n\nChamberlain Care® is the guiding principle we provide to each other and to our students to help our students achieve their goals, reach their dreams and impact healthcare at a national level.\n\nIf you're ready to live Chamberlain Care®, we invite you to apply today!\n\nAdtalem believes in a culture of collaboration, creativity, and innovation, which thrives when we can work flexibly. Depending on your location, this role may be hybrid or remote.\n\nHYBRID – If you reside within 50 miles of Adtalem offices in Chicago, IL, Addison, IL, Columbia, MD, Minneapolis, MN, and Miramar, FL this role is hybrid, requiring office presence four days a week (Mon, Tue, Wed, Thu), with the rest of the workweek spent working from home.\n\nJob Description\n\nOpportunity at a Glance\n\nThe Clinical Learning Developer is responsible for collaborating with leaders, faculty, and Curriculum and Instruction Developers to design and implement innovative, high quality clinical and simulation based learning content and experiences within the pre-licensure curriculum that help students apply knowledge and achieve clinical learning competencies, as well as, course and program outcomes. The process of designing relevant clinical and simulation experiences will include creation/development, process improvement, teaching/learning strategies, and assessment/evaluation.\n\nResponsibilities\n\nDesigns and integrates clinical and simulation learning into pre-licensure nursing courses\nResearches best-in-class clinical and simulation learning practices and incorporates them into courses being developed or revised\nUses data regarding course relevance, effectiveness and quality to make improvements\nIn partnership with the Curriculum and Instruction Developers, facilitates standardization of clinical and simulation based experiences for pre-licensure nursing courses.\nCollaborates with national academic and operational leadership as well as campus and online leadership to assess needs; recommend policies and procedures and performance indicators to facilitate adequate resources are available for excellent clinical and simulation based experiences for students.\nDevelops national clinical learning lab operating policies, procedures, and best practices throughout Chamberlain College of Nursing.\nProvides on-going assessment and direction to colleagues for standard operating procedures and systems used to support excellence in clinical and simulation based experiences for Chamberlain students.\nPerforms other duties as assigned\nComplies with all policies and standards\n\nQualifications\n\nMaster's Degree in Nursing Required\n5 experience in a clinical setting Preferred and\n3 simulation experience Preferred and\n5 working in a clinical healthcare environment Preferred and\nPrevious curriculum/course development experience Preferred and\nExperience with audio/visual technology. Required and\nExperience in adult education and/or training Preferred\nProficiency with Microsoft Office applications\nExcellent communication skills\nStrong analytical and problem-solving skills to resolve issues and determine effective business and alternative solutions.\nUnderstanding of the regulatory environment of higher education and nursing\nAbility to interact with all levels within the organization and work in a fast-paced environment with many initiatives.\nDemonstrated competence in leading successful meetings and projects.\nAbility to work effectively and productively in a remote location, including virtual meetings\nAbility to produce high quality outcomes related to multiple initiatives simultaneously\nAbility to function as a productive team member and to consistently model the TEACH values\nCertified Nurse Educator designation Preferred\n\nAdditional Information\n\nIn support of the pay transparency laws enacted across the country, the expected salary range for this position is between $70696.34 and $124230.14. Actual pay will be adjusted based on job-related factors permitted by law, such as experience and training; geographic location; licensure and certifications; market factors; departmental budgets; and responsibility. Our Talent Acquisition Team will be happy to answer any questions you may have, and we look forward to learning more about your salary requirements. The position qualifies for the below benefits.\n\nAdtalem Offers a Robust Suite Of Benefits Including\n\nHealth, dental, vision, life and disability insurance\n401k Retirement Program + 6% employer match\nParticipation in Adtalem’s Flexible Time Off (FTO) Policy\n12 Paid Holidays\n\nFor more information related to our benefits please visit: https://careers.adtalem.com/benefits. \n\nEqual Opportunity – Minority / Female / Disability / V / Gender Identity / Sexual Orientation",
    "company_description": "Adtalem Global Education (NYSE: ATGE), a leading medical and healthcare education provider, partners with organizations in the healthcare industry to solve critical workforce talent needs by expanding access to education, certifications and upskilling programs at scale.\n\nWith a dedicated focus on driving strong outcomes that increase workforce preparedness, Adtalem empowers a diverse learner population to achieve their goals and make inspiring contributions to the global community. Adtalem is the parent organization of American University of the Caribbean School of Medicine, Chamberlain University, Ross University School of Medicine, Ross University School of Veterinary Medicine and Walden University.\n\nAdtalem has more than 10,000 employees, a network of more than 275,000 alumni.\n\nWe believe diversity, equity and inclusion start at the top and must be intentional to be effective. Our commitment to DE&I, Sustainability and Corporate Governance have earned us the following recognitions:\n\n- Named one of Inc.'s Best Led Companies 2021\n- Ranked by Forbes as one of America’s Best Employers for Diversity 2021 & 2022\n- Ranked by Forbes as one of America’s Best Employers by State 2021\n- Recognized by Newsweek as one of America’s Most Responsible Companies 2021 & 2023\n- Named a Gold Award Winner 2021 in the Inspiring Diversity and Inclusion category by Inspiring Workplaces.\n- Named a Top Employer in Higher Education 2021 by DiversityJobs.Com\n…show more",
    "applicants": null,
  },
  "platform_url": "https://www.linkedin.com/jobs/view/4048357252/"
}
"""
            },
        ],
        response_format=JobPost
    )

    content = completion.choices[0].message.content  # Accessing the content correctly

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