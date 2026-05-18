const apiUrl = '/api/paciente';
const listUrl = '/api/pacientes';

// Elementos do DOM
const tbody = document.getElementById('pacientesTbody');
const modal = document.getElementById('pacienteModal');
const form = document.getElementById('pacienteForm');
const modalTitle = document.getElementById('modalTitle');
const searchInput = document.getElementById('searchInput');

// Estado
let pacientesData = [];

// Funções de Inicialização
document.addEventListener('DOMContentLoaded', carregarPacientes);

searchInput.addEventListener('input', (e) => {
    const term = e.target.value.toLowerCase();
    const filtrados = pacientesData.filter(p => 
        p.nome.toLowerCase().includes(term) || 
        p.cpf.includes(term) ||
        p.diagnostico.toLowerCase().includes(term)
    );
    renderTable(filtrados);
});

// Carregar pacientes da API
async function carregarPacientes() {
    try {
        const res = await fetch(listUrl);
        if (!res.ok) throw new Error('Erro ao carregar dados');
        const data = await res.json();
        pacientesData = data || [];
        renderTable(pacientesData);
    } catch (error) {
        console.error(error);
        alert('Não foi possível carregar os pacientes.');
    }
}

// Renderizar tabela
function renderTable(pacientes) {
    tbody.innerHTML = '';
    if (pacientes.length === 0) {
        tbody.innerHTML = '<tr><td colspan="7" style="text-align: center;">Nenhum paciente encontrado.</td></tr>';
        return;
    }
    
    pacientes.forEach(p => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>#${p.id}</td>
            <td><strong>${p.nome}</strong><br><small style="color: #64748b">${p.email}</small></td>
            <td>${p.cpf}</td>
            <td>${p.idade} anos</td>
            <td>${p.telefone}</td>
            <td>${p.diagnostico || '-'}</td>
            <td>
                <button class="btn-small btn-edit" onclick='editPaciente(${JSON.stringify(p)})'>Editar</button>
                <button class="btn-small btn-danger" onclick="deletePaciente(${p.id})">Excluir</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

// Modal Control
function openModal() {
    form.reset();
    document.getElementById('pacienteId').value = '';
    modalTitle.textContent = 'Cadastrar Paciente';
    modal.classList.add('active');
}

function closeModal() {
    modal.classList.remove('active');
}

// Submissão do Formulário
form.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const id = document.getElementById('pacienteId').value;
    const isEdit = !!id;
    
    const paciente = {
        nome: document.getElementById('nome').value,
        email: document.getElementById('email').value,
        cpf: document.getElementById('cpf').value,
        idade: parseInt(document.getElementById('idade').value, 10),
        telefone: document.getElementById('telefone').value,
        diagnostico: document.getElementById('diagnostico').value
    };

    if (isEdit) {
        paciente.id = parseInt(id, 10);
    }

    const endpoint = isEdit ? `${apiUrl}/update` : `${apiUrl}/create`;
    const method = isEdit ? 'PUT' : 'POST';

    try {
        const res = await fetch(endpoint, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(paciente)
        });

        if (!res.ok) throw new Error('Erro ao salvar paciente');
        
        closeModal();
        carregarPacientes();
    } catch (error) {
        console.error(error);
        alert('Erro ao salvar os dados.');
    }
});

// Editar Paciente
window.editPaciente = function(paciente) {
    document.getElementById('pacienteId').value = paciente.id;
    document.getElementById('nome').value = paciente.nome;
    document.getElementById('email').value = paciente.email;
    document.getElementById('cpf').value = paciente.cpf;
    document.getElementById('idade').value = paciente.idade;
    document.getElementById('telefone').value = paciente.telefone;
    document.getElementById('diagnostico').value = paciente.diagnostico;
    
    modalTitle.textContent = 'Editar Paciente';
    modal.classList.add('active');
};

// Excluir Paciente
window.deletePaciente = async function(id) {
    if (!confirm('Tem certeza que deseja excluir este paciente?')) return;

    try {
        const res = await fetch(`${apiUrl}/delete?id=${id}`, {
            method: 'DELETE'
        });
        
        if (!res.ok) throw new Error('Erro ao excluir');
        carregarPacientes();
    } catch (error) {
        console.error(error);
        alert('Erro ao excluir paciente.');
    }
};
