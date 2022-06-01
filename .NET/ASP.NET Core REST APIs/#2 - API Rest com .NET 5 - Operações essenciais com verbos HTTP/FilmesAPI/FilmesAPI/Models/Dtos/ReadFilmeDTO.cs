using System;
using System.ComponentModel.DataAnnotations;

namespace FilmesAPI.Models
{
    public class ReadFilmeDTO : FilmeDTO
    {
        [Key][Required]
        public int Id { get; set; }
        public DateTime HoraDaConsulta { get; set; }
    }
}
