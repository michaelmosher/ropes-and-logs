# Ropes and Logs

Building my own Raft implementation

## Essential Aspects

- Leader election
- Log replication
    - Important but non-essential: log compaction
- Safety
- Membership changes

> Raft implements consensus by first electing a distinguished leader, then 
> giving the leader complete responsibility for managing the replicated log.
> The leader accepts log entries from clients, replicates them on other
> servers, and tells servers when it is safe to apply log entries to their
> state machines.

### Notes

- "Replicated" means sent to followers; a change is "committed" by a leader
  once it knows a majority of servers have successfully applied it; a change is
  "committed" by a follower when a leader tells them it has been committed.

- broadcast time = average(network roundtrip time + write to storage time).
- election timeout should be at least 10x broadcast time (typically 10-500ms).

```mermaid
classDiagram

%% log: "each entry contains command for state machine, and term when entry was
%% received by leader (first index is 1)."
class LogEntry {
    int term
    command
}

LogEntry --> AppendEntriesInput

class AppendEntriesInput {
    int term
    string leaderId
    int prevLogIndex
    int prevLogTerm
    LogEntry[] entries
    int leaderCommit
}

class RequestVoteInput {
    string candidateId
    int term
    int lastLogIndex
    int lastLogTerm
}

LogEntry --> DataRepository

class DataRepository {
    int currentTerm
    int? votedFor
    LogEntry[] log
    
    updateTerm(int)
    updateVotedFor(int)
    appendEntries(LogEntry[])
}

DataRepository --> RaftServer
RequestVoteInput --> RaftServer
AppendEntriesInput --> RaftServer

%% currentTerm, votedFor, and log must be persistent (ie. "Updated on stable
%% storage before responding to RPCs").
%% election timeout? (generated with some randomness)
class RaftServer {
    string serverId
    DataRepository data
    int commitIndex
    int lastApplied

    AppendEntries(AppendEntriesInput) currentTerm success
    RequestVote(RequestVoteInput) term voteGranted
}

RaftServer --> RaftFollower

class RaftFollower {
    convertToCandidate()
}

RaftServer --> RaftCandidate

class RaftCandidate {
    map[string]bool voteReceived

    requestVote(RequestVoteInput) term voteGranted
    convertToFollower()
    convertToLeader()
}

RaftServer --> RaftLeader 

class RaftLeader {
    map[string]int nextIndex
    map[string]int matchIndex

    appendEntries(AppendEntriesInput) currentTerm success
    convertToFollower()
    ExecuteCommand(command) success
}

```

### Server States

```mermaid
stateDiagram-v2
direction TB

[*] --> Follower: starts up
Follower --> Candidate: times out, starts election
Candidate --> Candidate: times out, start new election
Candidate --> Leader: receives votes from majority of servers
Candidate --> Follower: discovers current leader or new term
Leader --> Follower: discovers server with higher term
```
