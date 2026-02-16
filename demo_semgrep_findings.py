# demo_semgrep_findings.py
import pickle
import subprocess

def unsafe_deserialize(user_bytes: bytes):
    # Finding #1: unsafe deserialization (pickle.loads)
    return pickle.loads(user_bytes)

def unsafe_shell(cmd: str):
    # Finding #2: command injection risk (shell=True)
    return subprocess.run(cmd, shell=True, check=False)
