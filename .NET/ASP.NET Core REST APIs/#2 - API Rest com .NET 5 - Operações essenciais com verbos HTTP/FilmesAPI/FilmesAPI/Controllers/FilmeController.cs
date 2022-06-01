using System;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;
using FilmesAPI.Models;
using System.Linq;
using FilmesAPI.Data;
using AutoMapper;

namespace FilmesAPI.Controllers
{
    [ApiController]
    [Route("[controller]")]

    public class FilmeController : ControllerBase
    {
        private FilmeContext _context;
        private IMapper _mapper;

        public FilmeController(FilmeContext context, IMapper mapper)
        {
            _context = context;
            _mapper = mapper;
        }

        [HttpPost]
        public IActionResult AdicionaFilme([FromBody] CreateFilmeDTO filmedto)
        {
            Filme filme = _mapper.Map<Filme>(filmedto);

            _context.Filmes.Add(filme);
            _context.SaveChanges();
            return CreatedAtAction(nameof(RecuperaFilmeByID), new { id = filme.id }, filme);
        }

        [HttpGet]
        public IActionResult RecuperaFilmes()
        {
            return Ok(_context.Filmes);
        }

        [HttpGet("{id}")]
        public IActionResult RecuperaFilmeByID(int id)
        {
            Filme filme = _context.Filmes.FirstOrDefault(filme => filme.id == id);
            if (filme != null)
            {
                ReadFilmeDTO readFilmeDto = _mapper.Map<ReadFilmeDTO>(filme);
                readFilmeDto.HoraDaConsulta = DateTime.Now;
                return Ok(readFilmeDto);
            }
            return NotFound();
            
        }

        [HttpPut("{id}")]
        public IActionResult AtualizaFilme(int id, [FromBody] UpdateFilmeDTO filmeAtualizado)
        {
            Filme filme = _context.Filmes.FirstOrDefault(filme => filme.id == id);
            if(filme != null)
            {
                _mapper.Map(filmeAtualizado, filme);
                _context.SaveChanges();
                return NoContent();
            }
            return NotFound();
        }

        [HttpDelete ("{id}")]
        public IActionResult DeletaFilme(int id)
        {
            Filme filme = _context.Filmes.FirstOrDefault(filme => filme.id == id);
            if(filme != null)
            {
                _context.Filmes.Remove(filme);
                _context.SaveChanges();
                return NoContent(); 
            }
            return NotFound();
        }
    }
}
