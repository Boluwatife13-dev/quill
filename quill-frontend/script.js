
const $ = id => document.getElementById(id);

let notes = JSON.parse(localStorage.getItem("quillNotes")) || [];
let tasks = JSON.parse(localStorage.getItem("quillTasks")) || [];
let currentNote = null;
let taskFilter = "all";


/* AUTH */

$("showSignup").onclick = e => {
    e.preventDefault();
    $("loginForm").classList.add("hidden");
    $("signupForm").classList.remove("hidden");
};

$("showLogin").onclick = e => {
    e.preventDefault();
    $("signupForm").classList.add("hidden");
    $("loginForm").classList.remove("hidden");
};

$("signupForm").onsubmit = e => {
    e.preventDefault();

    localStorage.setItem("quillUser", JSON.stringify({
        name: $("signupName").value,
        email: $("signupEmail").value,
        password: $("signupPassword").value
    }));

    $("signupMessage").textContent = "Account created. You can now login.";
    $("signupForm").reset();
};

$("loginForm").onsubmit = e => {
    e.preventDefault();

    const user = JSON.parse(localStorage.getItem("quillUser"));

    if (
        user &&
        $("loginEmail").value === user.email &&
        $("loginPassword").value === user.password
    ) {
        $("auth").classList.add("hidden");
        $("app").classList.remove("hidden");
        renderNotes();
        renderTasks();
    } else {
        $("loginMessage").textContent = "Invalid email or password.";
    }
};

$("logoutBtn").onclick = () => {
    $("app").classList.add("hidden");
    $("auth").classList.remove("hidden");
};


/* NAVIGATION */

$("notesTab").onclick = () => {
    $("editor").classList.remove("hidden");
    $("todo").classList.add("hidden");
};

$("tasksTab").onclick = () => {
    $("editor").classList.add("hidden");
    $("todo").classList.remove("hidden");
};


/* NOTES */

function renderNotes() {
    $("noteItems").innerHTML = "";

    notes.forEach(note => {
        const item = document.createElement("button");
        item.className = "note-item";
        item.textContent = note.title || "Untitled Note";

        item.onclick = () => openNote(note.id);

        $("noteItems").appendChild(item);
    });

    $("noteCount").textContent = notes.length;
}

function openNote(id) {
    const note = notes.find(n => n.id === id);
    if (!note) return;

    currentNote = id;
    $("noteTitle").value = note.title;
    $("noteContent").value = note.content;
    $("status").textContent = "● Saved";
    updateCounter();
}

$("newNoteBtn").onclick = () => {
    currentNote = null;
    $("noteTitle").value = "";
    $("noteContent").value = "";
    $("status").textContent = "● Unsaved";
};

$("saveBtn").onclick = () => {
    const title = $("noteTitle").value.trim() || "Untitled Note";
    const content = $("noteContent").value;

    if (currentNote === null) {
        currentNote = Date.now();

        notes.push({
            id: currentNote,
            title: title,
            content: content,
            pinned: false,
            updatedAt: Date.now()
        });
    } else {
        const note = notes.find(n => n.id === currentNote);

        if (note) {
            note.title = title;
            note.content = content;
            note.updatedAt = Date.now();
        }
    }

    localStorage.setItem("quillNotes", JSON.stringify(notes));

    $("status").textContent = "● Saved";

    renderNotes();
    updateCounter();
};

$("pinBtn").onclick = () => {
    if (!currentNote) return;

    const note = notes.find(n => n.id === currentNote);
    note.pinned = !note.pinned;

    $("pinBtn").textContent = note.pinned ? "★" : "☆";
    localStorage.setItem("quillNotes", JSON.stringify(notes));
};

$("noteContent").oninput = () => {
    $("status").textContent = "● Unsaved";
    updateCounter();
};

$("noteTitle").oninput = () => {
    $("status").textContent = "● Unsaved";
};

function updateCounter() {
    const text = $("noteContent").value.trim();

    $("words").textContent =
        "Words: " + (text ? text.split(/\s+/).length : 0);

    $("chars").textContent =
        "Characters: " + $("noteContent").value.length;
}


/* NOTE SEARCH */

$("searchNotes").oninput = e => {
    const value = e.target.value.toLowerCase();

    document.querySelectorAll(".note-item").forEach((item, i) => {
        const note = notes[i];
        item.style.display =
            note.title.toLowerCase().includes(value) ? "block" : "none";
    });
};

$("pinnedBtn").onclick = () => {
    document.querySelectorAll(".note-item").forEach((item, i) => {
        item.style.display = notes[i].pinned ? "block" : "none";
    });
};

$("allNotesBtn").onclick = renderNotes;

$("recentBtn").onclick = () => {
    [...notes]
        .sort((a,b) => b.updatedAt - a.updatedAt)
        .forEach(() => {});
    renderNotes();
};


/* TODO */

$("addTaskBtn").onclick = () => {
    $("taskModal").classList.remove("hidden");
};

$("closeModal").onclick =
$("cancelTask").onclick = () => {
    $("taskModal").classList.add("hidden");
};

$("taskForm").onsubmit = e => {
    e.preventDefault();

    tasks.push({
        id: Date.now(),
        title: $("taskTitle").value,
        description: $("taskDescription").value,
        priority: $("taskPriority").value,
        dueDate: $("taskDate").value,
        category: $("taskCategory").value,
        completed: false
    });

    localStorage.setItem("quillTasks", JSON.stringify(tasks));

    $("taskForm").reset();
    $("taskModal").classList.add("hidden");
    renderTasks();
};

function renderTasks() {
    $("taskItems").innerHTML = "";

    const filtered = tasks.filter(task =>
        taskFilter === "all" ||
        (taskFilter === "active" && !task.completed) ||
        (taskFilter === "completed" && task.completed)
    );

    filtered.forEach(task => {
        const div = document.createElement("div");
        div.className = "task" + (task.completed ? " completed" : "");

        div.innerHTML = `
            <input type="checkbox" ${task.completed ? "checked" : ""}>
            <div class="task-info">
                <div class="task-title">${task.title}</div>
                <small>${task.priority} ${task.category ? "• " + task.category : ""}</small>
            </div>
            <button>Delete</button>
        `;

        div.querySelector("input").onchange = () => {
            task.completed = !task.completed;
            localStorage.setItem("quillTasks", JSON.stringify(tasks));
            renderTasks();
        };

        div.querySelector("button").onclick = () => {
            tasks = tasks.filter(t => t.id !== task.id);
            localStorage.setItem("quillTasks", JSON.stringify(tasks));
            renderTasks();
        };

        $("taskItems").appendChild(div);
    });

    $("taskCount").textContent =
        filtered.length + (filtered.length === 1 ? " task" : " tasks");
}

document.querySelectorAll(".filters button").forEach(button => {
    button.onclick = () => {
        document.querySelectorAll(".filters button")
            .forEach(b => b.classList.remove("active"));

        button.classList.add("active");
        taskFilter = button.dataset.filter;
        renderTasks();
    };
});


/* DARK MODE */

$("themeBtn").onclick = () => {
    document.body.classList.toggle("dark");
};

