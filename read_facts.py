import pandas as pd
import os
import glob
from dotenv import load_dotenv
import psycopg2
from sqlalchemy import create_engine, text

def get_db_connection():
    load_dotenv()
    
    db_url = os.getenv('DB_CONNECTION_STRING')
    if not db_url:
        raise ValueError("DB_CONNECTION_STRING not found in environment variables")
    
    return create_engine(db_url)

csv_files = glob.glob('facts/*.csv')

all_facts = []
all_species = []

for file_path in csv_files:
    species = os.path.basename(file_path).replace('.csv', '')
    
    with open(file_path, 'r') as file:
        content = file.read()
        facts = [fact.strip() for fact in content.split('|')]
        
        all_facts.extend(facts)
        all_species.extend([species] * len(facts))

df = pd.DataFrame({
    'fact': all_facts,
    'species': all_species
})

try:
    engine = get_db_connection()
    
    with engine.connect() as conn:
        create_table_sql = text("""
            CREATE TABLE IF NOT EXISTS facts (
                id SERIAL PRIMARY KEY,
                fact TEXT NOT NULL,
                species VARCHAR(255) NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            )
        """)
        conn.execute(create_table_sql)
        conn.commit()
    
    df.to_sql('facts', engine, if_exists='append', index=False)
    print(f"Successfully wrote {len(df)} records to the database")
    
except Exception as e:
    print(f"Error writing to database: {str(e)}")
finally:
    if 'engine' in locals():
        engine.dispose()

print("\nDataFrame Shape:", df.shape)
print("\nFirst few rows of the DataFrame:")
print(df.head())
print("\nSample of facts for each species:")
print(df.groupby('species').size()) 